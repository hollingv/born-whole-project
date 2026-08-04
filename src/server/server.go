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

const aiModel = "@cf/meta/llama-3.1-8b-instruct"

const systemPrompt = `You are a helpful assistant answering questions about circumcision,
bodily autonomy, and children's rights. Answer based only on the provided context.
Be concise, factual, and compassionate. If the context does not contain enough
information to answer, say so.`

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
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if accountID == "" || apiToken == "" {
		return "", fmt.Errorf("CLOUDFLARE_ACCOUNT_ID and CLOUDFLARE_API_TOKEN env vars not set")
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

		// Keyword search for relevant chunks
		type chunk struct{ text, source string }
		var chunks []chunk
		qLower := strings.ToLower(q)
		for _, name := range filenames {
			content, err := os.ReadFile(filepath.Join(siteDir, "kb", name))
			if err != nil {
				continue
			}
			source := filenameToSource(name)
			for _, line := range strings.Split(string(content), "\n") {
				line = strings.TrimSpace(line)
				if line != "" && strings.Contains(strings.ToLower(line), qLower) {
					chunks = append(chunks, chunk{text: line, source: source})
				}
				if len(chunks) >= 10 {
					break
				}
			}
			if len(chunks) >= 10 {
				break
			}
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
