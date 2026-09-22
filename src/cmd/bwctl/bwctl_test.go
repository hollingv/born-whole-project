package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"born-whole-project/src/cmd/bwctl/internal/data"
)

func TestKBSourceURLs(t *testing.T) {
	// Verify that every static KB file in site/kb/ has a valid # Source: header
	// and that the URL in that header matches the Website field of a known
	// organization in organizations.go. This prevents the source URL shown in
	// AI responses from drifting out of sync with the organization data.

	// Build set of known org website URLs for fast lookup
	knownURLs := map[string]bool{}
	for _, group := range data.OrgGroups {
		for _, org := range group.Organizations {
			knownURLs[strings.TrimRight(org.Website, "/")] = true
		}
	}

	files, err := filepath.Glob("../../../site/kb/*.txt")
	if err != nil {
		t.Fatalf("could not glob kb directory: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("no .txt files found in site/kb/")
	}

	for _, path := range files {
		base := filepath.Base(path)
		f, err := os.Open(path)
		if err != nil {
			t.Errorf("%s: could not open file: %v", base, err)
			continue
		}
		defer f.Close()

		var sourceURL string
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if strings.HasPrefix(line, "# Source:") {
				sourceURL = strings.TrimSpace(strings.TrimPrefix(line, "# Source:"))
				break
			}
		}

		if sourceURL == "" {
			t.Errorf("%s: missing # Source: header", base)
			continue
		}

		normalized := strings.TrimRight(sourceURL, "/")
		if !knownURLs[normalized] {
			t.Errorf("%s: source URL %q does not match any organization Website", base, sourceURL)
		}
	}
}

func TestEachGroupHasOrganizations(t *testing.T) {
	for _, group := range data.OrgGroups {
		if len(group.Organizations) == 0 {
			t.Errorf("group %q has no organizations", group.Type)
		}
	}
}
