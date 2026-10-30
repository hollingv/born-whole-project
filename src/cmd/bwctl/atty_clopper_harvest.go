package main

import (
	"fmt"
	"os"
	"path/filepath"

	"born-whole-project/src/cmd/bwctl/internal/data"
)

const attyClopperKBDir = "site/kb/attyClopper"

// fetchNextTranscript finds the first Attorney Clopper YouTube Short that
// does not yet have a transcript file in site/kb/attyClopper/ and fetches it from
// the TranscribeYouTube API. Only one video is processed per run so that the
// caller controls the rate manually. Run the command repeatedly to accumulate
// the full transcript set.
func fetchNextTranscript() {
	resources := data.LoadResources()
	if len(resources) == 0 {
		fmt.Fprintf(os.Stderr, "[ WARN ] No YouTube Shorts metadata found in %s.\n", data.ResourcesPath)
		fmt.Fprintf(os.Stderr, "         Run 'bwctl harvest --source youtubeshorts' first.\n")
		return
	}

	if err := os.MkdirAll(attyClopperKBDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "[ ERROR ] Could not create %s: %v\n", attyClopperKBDir, err)
		return
	}

	for _, resource := range resources {
		outPath := filepath.Join(attyClopperKBDir, resource.VideoID+".txt")
		if _, err := os.Stat(outPath); err == nil {
			continue // already present — skip
		}

		fmt.Printf("Fetching: %s — %s\n", resource.VideoID, resource.Title)
		transcript, err := fetchTranscript(resource.VideoID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ SKIP ] %s: %v\n", resource.VideoID, err)
			return
		}

		content := fmt.Sprintf("# Source: https://www.youtube.com/shorts/%s\n# Title: %s\n\n%s\n",
			resource.VideoID, resource.Title, transcript)
		if err := os.WriteFile(outPath, []byte(content), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "[ WARN ] Could not write %s: %v\n", outPath, err)
			return
		}

		fmt.Printf("Written: %s\n", outPath)
		return // one at a time — done
	}

	fmt.Println("All transcripts already present in site/kb/attyClopper/")
}
