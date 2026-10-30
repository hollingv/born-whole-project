// Package main implements the local development server for Born Whole Project.
// It serves the generated site from the dist/ directory and provides the /ask
// and /ask-attorney-clopper endpoints, each searching a different KB manifest
// and calling the Cloudflare AI API to generate grounded answers.
// The handler logic mirrors functions/ask.js and functions/ask-attorney-clopper.js
// — keep both in sync when making changes.
package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// sentenceRe matches a single sentence ending with . ! or ?
var sentenceRe = regexp.MustCompile(`[^.!?]*[.!?]+\s*`)

// formatParagraphs splits plain AI response text into <p> elements by grouping
// every two sentences, regardless of how the model chose to format its output.
func formatParagraphs(text string) string {
	sentences := sentenceRe.FindAllString(text, -1)
	if len(sentences) == 0 {
		return "<p>" + text + "</p>"
	}
	var paras []string
	for i := 0; i < len(sentences); i += 2 {
		end := i + 2
		if end > len(sentences) {
			end = len(sentences)
		}
		if para := strings.TrimSpace(strings.Join(sentences[i:end], "")); para != "" {
			paras = append(paras, "<p>"+para+"</p>")
		}
	}
	return strings.Join(paras, "\n")
}

// uniqueSources returns deduplicated source names from a slice of chunks.
func uniqueSources(chunks []chunk) []string {
	seen := map[string]bool{}
	var sources []string
	for _, c := range chunks {
		if !seen[c.source] {
			seen[c.source] = true
			sources = append(sources, c.source)
		}
	}
	return sources
}

// handleAsk is the generic ask handler. It searches the given KB manifest and
// calls the Cloudflare AI API to produce a grounded answer.
func handleAsk(w http.ResponseWriter, r *http.Request, siteDir, manifestFile, disclaimer string) {
	q := r.URL.Query().Get("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	chunks, err := searchKB(q, siteDir, manifestFile)
	if err != nil {
		fmt.Fprintf(w, "<p>%s</p>", err.Error())
		return
	}
	if len(chunks) == 0 {
		fmt.Fprintf(w, "<p>No information found for: <strong>%s</strong></p>", q)
		return
	}

	sources := uniqueSources(chunks)
	var texts []string
	for _, c := range chunks {
		texts = append(texts, c.text)
	}

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

	fmt.Fprintf(w, "%s", formatParagraphs(answer))
	var sourceLinks []string
	for _, s := range sources {
		sourceLinks = append(sourceLinks, fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, s, s))
	}
	fmt.Fprintf(w, "<p class=\"ask-sources\">Sources:<br>%s</p>", strings.Join(sourceLinks, "<br>"))
	fmt.Fprintf(w, "<p class=\"ask-disclaimer\">%s</p>", disclaimer)
}

func newHandler(siteDir string) http.Handler {
	fs := http.FileServer(http.Dir(siteDir))
	mux := http.NewServeMux()

	mux.HandleFunc("/ask", func(w http.ResponseWriter, r *http.Request) {
		handleAsk(w, r, siteDir,
			"kb-bundle.json",
			"Answers are generated from curated sources. Always verify with the linked organisations.")
	})

	mux.HandleFunc("/ask-attorney-clopper", func(w http.ResponseWriter, r *http.Request) {
		handleAsk(w, r, siteDir,
			"attyclopper-kb-bundle.json",
			"Answers are drawn from Attorney Clopper's YouTube Shorts transcripts. Always verify with the source videos.")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
	return mux
}

func main() {
	log.Println("Serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", newHandler(targetDir)))
}
