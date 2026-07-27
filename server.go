package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// version is set at build time via -ldflags '-X main.version=...'.
// If not set (e.g. go run), it falls back to calling 'make app-tag'.
var version = ""

func init() {
	if version == "" {
		out, err := exec.Command("make", "app-tag").Output()
		if err == nil {
			version = strings.TrimSpace(string(out))
		} else {
			version = "unknown"
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("site"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			content, err := os.ReadFile("site/index.html")
			if err != nil {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			html := strings.ReplaceAll(string(content), "__COMMIT_SHA__", version)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write([]byte(html))
			return
		}
		fs.ServeHTTP(w, r)
	})

	log.Println("Serving on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
