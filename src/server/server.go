// Package main implements the local development server for Bodily Integrity Commons.
// The /ask handler logic mirrors functions/ask.js — keep both in sync when making changes.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const aiModel = "@cf/meta/llama-3.1-8b-instruct-fast"

const systemPrompt = `You are a helpful assistant answering questions about circumcision,
bodily autonomy, and children's rights. Only use the text provided below as context.
Do not use any outside knowledge. Be concise, factual, and compassionate.
If the provided context does not contain enough information to answer, say so explicitly.`

var stopWords = map[string]bool{
	"a": true, "an": true, "the": true, "is": true, "are": true, "was": true,
	"were": true, "be": true, "been": true, "being": true, "have": true, "has": true,
	"had": true, "do": true, "does": true, "did": true, "will": true, "would": true,
	"could": true, "should": true, "may": true, "might": true, "can": true,
	"what": true, "why": true, "how": true, "when": true, "where": true, "who": true,
	"which": true, "that": true, "this": true, "these": true, "those": true,
	"i": true, "me": true, "my": true, "we": true, "our": true, "you": true,
	"your": true, "it": true, "its": true, "they": true, "them": true, "their": true,
	"of": true, "in": true, "to": true, "for": true, "on": true, "with": true,
	"at": true, "by": true, "from": true, "about": true, "and": true, "or": true,
	"not": true, "no": true, "so": true, "if": true, "than": true, "very": true,
}

func extractKeywords(query string) []string {
	query = strings.ToLower(query)
	var words []string
	for _, word := range strings.FieldsFunc(query, func(r rune) bool {
		return !('a' <= r && r <= 'z')
	}) {
		if len(word) > 2 && !stopWords[word] {
			words = append(words, word)
		}
	}
	return words
}

// filenameToSource converts a KB filename to a readable source name.
func filenameToSource(filename string) string {
	name := strings.TrimSuffix(filepath.Base(filename), ".txt")
	if idx := strings.LastIndex(name, "-org"); idx != -1 {
		name = name[:idx] + ".org"
	}
	name = strings.ReplaceAll(name, "-", ".")
	return name
}

type cfAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type cfAIRequest struct {
	Messages []cfAIMessage `json:"messages"`
}

type cfAIResponse struct {
	Result struct {
		Response string `json:"response"`
	} `json:"result"`
	Success bool `json:"success"`
}

func callCloudflareAI(question string, chunks []string) (string, error) {
	accountID := os.Getenv(projectPrefix + "_CF_ACCOUNT_ID")
	apiToken := os.Getenv(projectPrefix + "_CF_API_TOKEN")
	if accountID == "" || apiToken == "" {
		return "", fmt.Errorf("%s_CF_ACCOUNT_ID and %s_CF_API_TOKEN env vars not set", projectPrefix, projectPrefix)
	}

	context := strings.Join(chunks, "\n\n")
	req := cfAIRequest{
		Messages: []cfAIMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: fmt.Sprintf("Context:\n%s\n\nQuestion: %s", context, question)},
		},
	}
	body, _ := json.Marshal(req)

	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/ai/run/%s", accountID, aiModel)
	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var aiResp cfAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&aiResp); err != nil {
		return "", err
	}
	if !aiResp.Success {
		return "", fmt.Errorf("Cloudflare AI returned success=false")
	}
	return aiResp.Result.Response, nil
}

func newHandler(siteDir string) http.Handler {
	fs := http.FileServer(http.Dir(siteDir))
	mux := http.NewServeMux()

	mux.HandleFunc("/ask", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Load KB manifest
		manifestData, err := os.ReadFile(filepath.Join(siteDir, "kb", "manifest.json"))
		if err != nil {
			fmt.Fprintf(w, "<p>Knowledge base not available.</p>")
			return
		}
		var filenames []string
		if err := json.Unmarshal(manifestData, &filenames); err != nil {
			fmt.Fprintf(w, "<p>Knowledge base error.</p>")
			return
		}

		// Extract keywords and search for relevant chunks
		type chunk struct {
			text, source string
			score        int
		}
		keywords := extractKeywords(q)
		if len(keywords) == 0 {
			keywords = []string{strings.ToLower(q)}
		}
		var chunks []chunk
		for _, name := range filenames {
			content, err := os.ReadFile(filepath.Join(siteDir, "kb", name))
			if err != nil {
				continue
			}
			source := filenameToSource(name)
			for _, para := range strings.Split(string(content), "\n\n") {
				para = strings.TrimSpace(para)
				if para == "" {
					continue
				}
				paraLower := strings.ToLower(para)
				score := 0
				for _, kw := range keywords {
					if strings.Contains(paraLower, kw) {
						score++
					}
				}
				if score > 0 {
					chunks = append(chunks, chunk{text: para, source: source, score: score})
				}
			}
		}
		// Sort by score descending and take top 10
		for i := 0; i < len(chunks)-1; i++ {
			for j := i + 1; j < len(chunks); j++ {
				if chunks[j].score > chunks[i].score {
					chunks[i], chunks[j] = chunks[j], chunks[i]
				}
			}
		}
		if len(chunks) > 10 {
			chunks = chunks[:10]
		}

		if len(chunks) == 0 {
			fmt.Fprintf(w, "<p>No information found for: <strong>%s</strong></p>", q)
			return
		}

		// Collect unique sources
		seen := map[string]bool{}
		var sources []string
		var texts []string
		for _, c := range chunks {
			texts = append(texts, c.text)
			if !seen[c.source] {
				seen[c.source] = true
				sources = append(sources, c.source)
			}
		}

		// Call Cloudflare AI
		answer, err := callCloudflareAI(q, texts)
		if err != nil {
			log.Printf("AI call failed: %v", err)
			fmt.Fprintf(w, "<p>AI unavailable. Here are relevant excerpts:</p><ul>")
			for _, c := range chunks {
				fmt.Fprintf(w, "<li>%s</li>", c.text)
			}
			fmt.Fprintf(w, "</ul>")
			return
		}

		fmt.Fprintf(w, "<p>%s</p>", answer)
		fmt.Fprintf(w, "<p class=\"ask-sources\">Sources: %s</p>", strings.Join(sources, ", "))
		fmt.Fprintf(w, "<p class=\"ask-disclaimer\">Answers are generated from curated sources. Always verify with the linked organisations.</p>")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})

	return mux
}

func main() {
	log.Println("Serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", newHandler("site")))
}
