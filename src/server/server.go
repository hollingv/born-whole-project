package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// filenameToSource converts a KB filename to a readable source name.
func filenameToSource(filename string) string {
	name := strings.TrimSuffix(filepath.Base(filename), ".txt")
	if idx := strings.LastIndex(name, "-org"); idx != -1 {
		name = name[:idx] + ".org"
	}
	name = strings.ReplaceAll(name, "-", ".")
	return name
}

func newHandler(siteDir string) http.Handler {
	fs := http.FileServer(http.Dir(siteDir))
	mux := http.NewServeMux()

	mux.HandleFunc("/ask", func(w http.ResponseWriter, r *http.Request) {
		q := strings.ToLower(r.URL.Query().Get("q"))
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		type match struct {
			source string
			line   string
		}
		var matches []match

		files, _ := filepath.Glob(filepath.Join(siteDir, "kb", "*.txt"))
		for _, f := range files {
			content, err := os.ReadFile(f)
			if err != nil {
				continue
			}
			source := filenameToSource(f)
			for _, line := range strings.Split(string(content), "\n") {
				line = strings.TrimSpace(line)
				if line != "" && strings.Contains(strings.ToLower(line), q) {
					matches = append(matches, match{source: source, line: line})
				}
				if len(matches) >= 10 {
					break
				}
			}
			if len(matches) >= 10 {
				break
			}
		}

		if len(matches) == 0 {
			fmt.Fprintf(w, "<p>No results found for: <strong>%s</strong></p>", r.URL.Query().Get("q"))
			return
		}

		fmt.Fprintf(w, "<p>Results for: <strong>%s</strong></p>", r.URL.Query().Get("q"))
		fmt.Fprintf(w, "<table class=\"ask-results\"><thead><tr><th>Source</th><th>Content</th></tr></thead><tbody>")
		for _, m := range matches {
			fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td></tr>", m.source, m.line)
		}
		fmt.Fprintf(w, "</tbody></table>")
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
