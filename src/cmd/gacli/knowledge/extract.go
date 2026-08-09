// Package knowledge provides HTML text extraction and paragraph normalisation
// utilities for building the site knowledge base.
package knowledge

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// ContentTags are the only HTML elements from which text is extracted.
var ContentTags = map[string]struct{}{
	"p": {}, "h1": {}, "h2": {}, "h3": {}, "h4": {}, "h5": {}, "h6": {},
	"article": {}, "main": {}, "section": {}, "blockquote": {}, "li": {},
}

// AllText recursively collects all text within a node and its descendants.
func AllText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(AllText(c))
	}
	return sb.String()
}

// ExtractText walks the HTML node tree and returns text only from known content elements.
func ExtractText(n *html.Node) string {
	if n.Type == html.ElementNode {
		if _, ok := ContentTags[n.Data]; ok {
			text := strings.TrimSpace(AllText(n))
			if text != "" {
				return text + "\n"
			}
			return ""
		}
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(ExtractText(c))
	}
	return sb.String()
}

// NormalizeToParagraphs groups lines of extracted text into coherent paragraphs
// separated by double newlines, filtering out very short fragments.
func NormalizeToParagraphs(text string) string {
	lines := strings.Split(text, "\n")
	var paragraphs []string
	var current []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			if len(current) > 0 {
				para := strings.Join(current, " ")
				if len(para) > 40 {
					paragraphs = append(paragraphs, para)
				}
				current = nil
			}
		} else {
			current = append(current, line)
		}
	}
	if len(current) > 0 {
		para := strings.Join(current, " ")
		if len(para) > 40 {
			paragraphs = append(paragraphs, para)
		}
	}
	return strings.Join(paragraphs, "\n\n")
}

// URLToFilename converts a URL into a safe filename.
func URLToFilename(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "unknown.txt"
	}
	name := strings.NewReplacer(".", "-", "/", "-").Replace(u.Hostname() + u.Path)
	name = strings.Trim(name, "-")
	return name + ".txt"
}
