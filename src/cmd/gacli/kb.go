package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

const kbDir = "site/kb"

var httpClient = &http.Client{Timeout: 15 * time.Second}

// normalizeToParagraphs groups lines of extracted text into coherent paragraphs
// separated by double newlines, filtering out very short fragments.
func normalizeToParagraphs(text string) string {
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

// contentTags are the only HTML elements from which text is extracted.
// This targets meaningful prose content and ignores navigation, UI, and boilerplate.
var contentTags = map[string]struct{}{
	"p": {}, "h1": {}, "h2": {}, "h3": {}, "h4": {}, "h5": {}, "h6": {},
	"article": {}, "main": {}, "section": {}, "blockquote": {}, "li": {},
}

// allText recursively collects all text within a node and its descendants.
func allText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(allText(c))
	}
	return sb.String()
}

// extractText walks the HTML node tree and returns text only from known content elements.
func extractText(n *html.Node) string {
	if n.Type == html.ElementNode {
		if _, ok := contentTags[n.Data]; ok {
			text := strings.TrimSpace(allText(n))
			if text != "" {
				return text + "\n"
			}
			return ""
		}
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(extractText(c))
	}
	return sb.String()
}

// urlToFilename converts a URL into a safe filename.
func urlToFilename(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "unknown.txt"
	}
	name := strings.NewReplacer(".", "-", "/", "-").Replace(u.Hostname() + u.Path)
	name = strings.Trim(name, "-")
	return name + ".txt"
}

var kbBuildCmd = &cobra.Command{
	Use:   "kb-build",
	Short: "Build the knowledge base from web sources",
	Long:  `Fetches each configured web source, extracts the text content, and saves it to the kb/ directory`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := os.RemoveAll(kbDir); err != nil {
			fmt.Fprintf(os.Stderr, "Error clearing kb directory: %v\n", err)
			os.Exit(1)
		}
		if err := os.MkdirAll(kbDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating kb directory: %v\n", err)
			os.Exit(1)
		}

		for _, group := range orgGroups {
			for _, org := range group.Organizations {
				urls := org.KBURLs
				if len(urls) == 0 {
					urls = []string{org.Website}
				}
				for _, u := range urls {
					fmt.Printf("Fetching %s (%s)...\n", org.Name, u)

					resp, err := httpClient.Get(u)
					if err != nil {
						fmt.Fprintf(os.Stderr, "  Error fetching %s: %v\n", u, err)
						continue
					}
					defer resp.Body.Close()
					if resp.StatusCode != http.StatusOK {
						fmt.Fprintf(os.Stderr, "  Skipping %s: HTTP %d\n", u, resp.StatusCode)
						continue
					}

					doc, err := html.Parse(resp.Body)
					if err != nil {
						fmt.Fprintf(os.Stderr, "  Error parsing HTML from %s: %v\n", u, err)
						continue
					}

					text := normalizeToParagraphs(extractText(doc))

					outPath := filepath.Join(kbDir, urlToFilename(u))
					if err := os.WriteFile(outPath, []byte(text), 0644); err != nil {
						fmt.Fprintf(os.Stderr, "  Error writing %s: %v\n", outPath, err)
						continue
					}

					fmt.Printf("  Saved to %s\n", outPath)
				}
			}
		}

		// Write manifest of all KB files for Cloudflare Pages Functions
		files, _ := filepath.Glob(filepath.Join(kbDir, "*.txt"))
		var filenames []string
		for _, f := range files {
			filenames = append(filenames, filepath.Base(f))
		}
		manifest, _ := json.MarshalIndent(filenames, "", "  ")
		manifestPath := filepath.Join(kbDir, "manifest.json")
		if err := os.WriteFile(manifestPath, manifest, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing manifest: %v\n", err)
		} else {
			fmt.Printf("Written %s\n", manifestPath)
		}

		fmt.Println("Knowledge base build complete.")
	},
}

func init() {
	rootCmd.AddCommand(kbBuildCmd)
}
