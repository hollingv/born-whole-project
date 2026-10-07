package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"born-whole-project/src/cmd/bwctl/internal/data"

	"golang.org/x/net/html"
)

const bsmEventsURL = "https://www.bloodstainedmen.com/about-us/events/"
const bsmOrgID = "blood-stained-men"

var eventsHTTPClient = &http.Client{Timeout: 15 * time.Second}

// findFirstHref returns the href of the first anchor tag within a node, or empty string.
func findFirstHref(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				return attr.Val
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if href := findFirstHref(c); href != "" {
			return href
		}
	}
	return ""
}

// tableRow holds the text cells and the href from the first cell of a table row.
type tableRow struct {
	cells []string
	href  string
}

// collectTableRows recursively collects all table row cell texts and hrefs from the HTML tree.
func collectTableRows(n *html.Node, rows *[]tableRow) {
	if n.Type == html.ElementNode && n.Data == "tr" {
		row := tableRow{}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode && (c.Data == "td" || c.Data == "th") {
				text := strings.TrimSpace(allText(c))
				if text != "" {
					row.cells = append(row.cells, text)
				}
				// Capture href from first cell only
				if len(row.cells) == 1 && row.href == "" {
					row.href = findFirstHref(c)
				}
			}
		}
		if len(row.cells) > 0 {
			*rows = append(*rows, row)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectTableRows(c, rows)
	}
}

// allText recursively collects all text within a node.
func allText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		chunk := allText(c)
		if chunk == "" {
			continue
		}
		if sb.Len() > 0 {
			last := sb.String()
			if last[len(last)-1] != ' ' && last[len(last)-1] != '\n' &&
				chunk[0] != ' ' && chunk[0] != '\n' {
				sb.WriteString(" ")
			}
		}
		sb.WriteString(chunk)
	}
	return sb.String()
}

// fetchAndParseBSMEvents fetches the Blood Stained Men events page and
// parses the table structure into a slice of Event entries.
func fetchAndParseBSMEvents() ([]data.Event, error) {
	fmt.Printf("Fetching Blood Stained Men events from %s...\n", bsmEventsURL)

	resp, err := eventsHTTPClient.Get(bsmEventsURL)
	if err != nil {
		return nil, fmt.Errorf("fetch BSM events: %w", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parse BSM events HTML: %w", err)
	}

	var rows []tableRow
	collectTableRows(doc, &rows)

	var events []data.Event
	var campaignTitle string
	var pendingEvent *data.Event

	for _, row := range rows {
		cells := row.cells
		first := strings.TrimSpace(cells[0])

		// Skip header rows
		if strings.EqualFold(first, "event page") || strings.EqualFold(first, "event") {
			continue
		}

		// Location row — always paired with the event row above
		if strings.HasPrefix(strings.ToLower(first), "location:") {
			if pendingEvent != nil {
				location := strings.TrimSpace(strings.TrimPrefix(strings.ToLower(first), "location:"))
				pendingEvent.Description += " | " + location
				events = append(events, *pendingEvent)
				pendingEvent = nil
			}
			continue
		}

		// Single-cell non-location row → campaign title
		if len(cells) == 1 {
			campaignTitle = first
			continue
		}

		// Event row: Day N | date | city | time
		if len(cells) >= 3 {
			if pendingEvent != nil {
				events = append(events, *pendingEvent)
			}
			desc := cells[2] // city
			if len(cells) >= 4 {
				desc += " — " + cells[3] // time
			}
			if campaignTitle != "" {
				desc = "[" + campaignTitle + "] " + desc
			}
			pendingEvent = &data.Event{
				Date:        cells[1],
				Description: desc,
				URL:         row.href,
				OrgID:       bsmOrgID,
			}
		}
	}
	if pendingEvent != nil {
		events = append(events, *pendingEvent)
	}

	fmt.Printf("  Parsed %d Blood Stained Men events\n", len(events))
	return events, nil
}

// harvestBSMEvents fetches BSM events and writes them to the harvested events JSON file.
func harvestBSMEvents() {
	events, err := fetchAndParseBSMEvents()
	if err != nil {
		fmt.Fprintf(os.Stderr, "  Skipping BSM events: %v\n", err)
		return
	}
	raw, _ := json.MarshalIndent(events, "", "  ")
	if err := os.WriteFile(data.HarvestedEventsPath, raw, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "  Error writing %s: %v\n", data.HarvestedEventsPath, err)
		return
	}
	fmt.Printf("  Written %s\n", data.HarvestedEventsPath)
}
