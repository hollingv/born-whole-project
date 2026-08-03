package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func newHandler(siteDir string) http.Handler {
	fs := http.FileServer(http.Dir(siteDir))
	mux := http.NewServeMux()

	mux.HandleFunc("/ask", func(w http.ResponseWriter, r *http.Request) {
		q := strings.ToLower(r.URL.Query().Get("q"))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		var matches []string
		files, _ := filepath.Glob(filepath.Join(siteDir, "kb", "*.txt"))
		for _, f := range files {
			content, err := os.ReadFile(f)
			if err != nil {
				continue
			}
			for _, line := range strings.Split(string(content), "\n") {
				line = strings.TrimSpace(line)
				if line != "" && strings.Contains(strings.ToLower(line), q) {
					matches = append(matches, line)
				}
				if len(matches) >= 5 {
					break
				}
			}
			if len(matches) >= 5 {
				break
			}
		}

		if len(matches) == 0 {
			fmt.Fprintf(w, "<p>No results found for: <strong>%s</strong></p>", r.URL.Query().Get("q"))
			return
		}
		fmt.Fprintf(w, "<p>Results for: <strong>%s</strong></p><ul>", r.URL.Query().Get("q"))
		for _, m := range matches {
			fmt.Fprintf(w, "<li>%s</li>", m)
		}
		fmt.Fprintf(w, "</ul>")
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
