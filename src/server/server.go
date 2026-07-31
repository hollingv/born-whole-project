package main

import (
	"fmt"
	"log"
	"net/http"
)

func newHandler(siteDir string) http.Handler {
	fs := http.FileServer(http.Dir(siteDir))
	mux := http.NewServeMux()

	mux.HandleFunc("/ask", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<p>You asked: <strong>%s</strong></p><p>Stub response: more information coming soon.</p>", q)
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
