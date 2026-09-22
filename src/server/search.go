package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var stopWords = map[string]struct{}{
	"a": {}, "an": {}, "the": {}, "is": {}, "are": {}, "was": {},
	"were": {}, "be": {}, "been": {}, "being": {}, "have": {}, "has": {},
	"had": {}, "do": {}, "does": {}, "did": {}, "will": {}, "would": {},
	"could": {}, "should": {}, "may": {}, "might": {}, "can": {},
	"what": {}, "why": {}, "how": {}, "when": {}, "where": {}, "who": {},
	"which": {}, "that": {}, "this": {}, "these": {}, "those": {},
	"i": {}, "me": {}, "my": {}, "we": {}, "our": {}, "you": {},
	"your": {}, "it": {}, "its": {}, "they": {}, "them": {}, "their": {},
	"of": {}, "in": {}, "to": {}, "for": {}, "on": {}, "with": {},
	"at": {}, "by": {}, "from": {}, "about": {}, "and": {}, "or": {},
	"not": {}, "no": {}, "so": {}, "if": {}, "than": {}, "very": {},
}

type chunk struct {
	text, source string
	score        int
}

// extractKeywords strips stop words and returns meaningful search terms.
func extractKeywords(query string) []string {
	query = strings.ToLower(query)
	var words []string
	for _, word := range strings.FieldsFunc(query, func(r rune) bool {
		return !('a' <= r && r <= 'z')
	}) {
		if _, isStop := stopWords[word]; len(word) > 2 && !isStop {
			words = append(words, word)
		}
	}
	return words
}

// sourceFromContent extracts the source URL from the "# Source: ..." header
// line of a KB file. Falls back to the filename if no header is found.
func sourceFromContent(content, filename string) string {
	for _, line := range strings.SplitN(content, "\n", 5) {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# Source:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# Source:"))
		}
	}
	return strings.TrimSuffix(filepath.Base(filename), ".txt")
}

// searchKB loads the KB manifest and returns scored, sorted chunks matching the query.
func searchKB(q, siteDir string) ([]chunk, error) {
	manifestData, err := os.ReadFile(filepath.Join(siteDir, "kb", "manifest.json"))
	if err != nil {
		return nil, fmt.Errorf("KB not available: %w", err)
	}
	var filenames []string
	if err := json.Unmarshal(manifestData, &filenames); err != nil {
		return nil, fmt.Errorf("KB manifest error: %w", err)
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
		source := sourceFromContent(string(content), name)
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
	return chunks, nil
}
