// Package main implements the local development server for Bodily Integrity Commons.
// The /ask handler logic mirrors functions/ask.js — keep both in sync when making changes.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

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

// handleAsk processes the /ask request — searches KB and calls AI.
func handleAsk(w http.ResponseWriter, r *http.Request, siteDir string) {
	q := r.URL.Query().Get("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	chunks, err := searchKB(q, siteDir)
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

	fmt.Fprintf(w, "<p>%s</p>", answer)
	var sourceLinks []string
	for _, s := range sources {
		sourceLinks = append(sourceLinks, fmt.Sprintf(`<a href="https://%s" target="_blank">%s</a>`, s, s))
	}
	fmt.Fprintf(w, "<p class=\"ask-sources\">Sources:<br>%s</p>", strings.Join(sourceLinks, "<br>"))
	fmt.Fprintf(w, "<p class=\"ask-disclaimer\">Answers are generated from curated sources. Always verify with the linked organisations.</p>")
}

func newHandler(siteDir string) http.Handler {
	fs := http.FileServer(http.Dir(siteDir))
	mux := http.NewServeMux()
	mux.HandleFunc("/ask", func(w http.ResponseWriter, r *http.Request) {
		handleAsk(w, r, siteDir)
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
