package main

import (
	"fmt"
	"os"
	"path/filepath"

	"born-whole-project/src/cmd/bwctl/internal/data"

	"github.com/spf13/cobra"
)

var ytShortsCmd = &cobra.Command{
	Use:   "yt-shorts",
	Short: "Manage Attorney Clopper YouTube Shorts transcripts for the AI knowledge base",
	Long: `Commands for building and inspecting the Attorney Clopper transcript
knowledge base in site/kb/attyClopper/. Transcripts are fetched one at a time
from the TranscribeYouTube API (transcribeyoutube.com) and committed to
git so they power AI queries without re-fetching.`,
}

var ytShortsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch the next missing transcript and write it to site/kb/attyClopper/",
	Long: `Finds the first YouTube Short in harvested-youtube-shorts.json that does
not yet have a transcript file in site/kb/attyClopper/, fetches its transcript
from the TranscribeYouTube API, writes the file, and exits.

Run once per request. Commit the result, then run again for the next video.

Rate limits: 5/min · 30/hr · 100/day per IP. A 429 is handled automatically
by sleeping the retryAfter duration before retrying the same video.`,
	Run: func(cmd *cobra.Command, args []string) {
		fetchNextTranscript()
	},
}

var ytShortsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List transcript status for every YouTube Short",
	Long:  `Prints a table showing which videos have been transcribed and which are still missing.`,
	Run: func(cmd *cobra.Command, args []string) {
		resources := data.LoadResources()
		if len(resources) == 0 {
			fmt.Fprintf(os.Stderr, "[ WARN ] No YouTube Shorts metadata found.\n")
			fmt.Fprintf(os.Stderr, "         Run 'bwctl harvest --source youtubeshorts' first.\n")
			return
		}

		fmt.Printf("%-15s  %-50s  %s\n", "VIDEO ID", "TITLE", "STATUS")
		fmt.Printf("%-15s  %-50s  %s\n", "---------------", "--------------------------------------------------", "------")

		var done, missing int
		for _, r := range resources {
			title := r.Title
			if len(title) > 50 {
				title = title[:47] + "..."
			}
			if _, err := os.Stat(filepath.Join(attyClopperKBDir, r.VideoID+".txt")); err == nil {
				fmt.Printf("%-15s  %-50s  ✓\n", r.VideoID, title)
				done++
			} else {
				fmt.Printf("%-15s  %-50s  ✗ missing\n", r.VideoID, title)
				missing++
			}
		}
		fmt.Printf("\n%d transcribed, %d missing, %d total\n", done, missing, done+missing)
	},
}

var ytShortsStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show a one-line summary of transcript progress",
	Run: func(cmd *cobra.Command, args []string) {
		resources := data.LoadResources()
		if len(resources) == 0 {
			fmt.Fprintf(os.Stderr, "[ WARN ] No YouTube Shorts metadata found.\n")
			fmt.Fprintf(os.Stderr, "         Run 'bwctl harvest --source youtubeshorts' first.\n")
			return
		}

		var done int
		for _, r := range resources {
			if _, err := os.Stat(filepath.Join(attyClopperKBDir, r.VideoID+".txt")); err == nil {
				done++
			}
		}
		total := len(resources)
		fmt.Printf("%d of %d transcribed (%d remaining)\n", done, total, total-done)
	},
}

func init() {
	ytShortsCmd.AddCommand(ytShortsGetCmd, ytShortsListCmd, ytShortsStatusCmd)
	rootCmd.AddCommand(ytShortsCmd)
}
