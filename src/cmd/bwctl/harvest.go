package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"born-whole-project/src/cmd/bwctl/internal/youtube"

	"github.com/spf13/cobra"
)

const kbDir = "site/kb"

// harvestedDir is where bwctl harvest writes its output files.
// These files are tracked by git so that a given tag always produces
// identical site content regardless of when the build runs.
const harvestedDir = "site/harvested"

const (
	sourceAll           = "all"
	sourceEvents        = "events"
	sourceYouTubeShorts = "youtubeshorts"
)

// writeManifest writes the list of KB text files to generated-kb-file-list.json
// in dist/harvested/. Called by bwctl site at build time, not by bwctl harvest,
// so the manifest is always in sync with whatever is on disk when the site is built.
// Paths are stored relative to dist/ (e.g. "kb/intact-global.txt").
func writeManifest() {
	var paths []string
	if files, _ := filepath.Glob(filepath.Join(kbDir, "*.txt")); files != nil {
		for _, f := range files {
			paths = append(paths, "kb/"+filepath.Base(f))
		}
	}
	if files, _ := filepath.Glob(filepath.Join(harvestedDir, "yt-*.txt")); files != nil {
		for _, f := range files {
			paths = append(paths, "harvested/"+filepath.Base(f))
		}
	}
	manifest, _ := json.MarshalIndent(paths, "", "  ")
	manifestPath := filepath.Join(targetDir, "harvested", "generated-kb-file-list.json")
	if err := os.WriteFile(manifestPath, manifest, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing manifest: %v\n", err)
		return
	}
	fmt.Printf("Written %s\n", manifestPath)
}

var harvestCmd = &cobra.Command{
	Use:   "harvest",
	Short: "Fetch live data into site/harvested/ for review and commit",
	Long: `Fetches live data and writes it to site/harvested/ so it can be
reviewed and committed to git. Run this manually when you want to
refresh events or YouTube Shorts metadata.
Use --source to select which data to harvest (default: all).`,
	Run: func(cmd *cobra.Command, args []string) {
		source, _ := cmd.Flags().GetString("source")
		days, _ := cmd.Flags().GetInt("days")

		switch source {
		case sourceAll, sourceEvents, sourceYouTubeShorts:
			// valid
		default:
			fmt.Fprintf(os.Stderr, "Invalid --source %q. Must be one of: %s, %s, or %s\n",
				source, sourceAll, sourceEvents, sourceYouTubeShorts)
			os.Exit(1)
		}

		if err := os.MkdirAll(harvestedDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating harvested directory: %v\n", err)
			os.Exit(1)
		}

		if source == sourceAll || source == sourceEvents {
			harvestBSMEvents()
		}
		if source == sourceAll || source == sourceYouTubeShorts {
			youtube.BuildResources(projectPrefix, days)
		}

		fmt.Println("Harvest complete. Review changes in site/harvested/ then commit.")
	},
}

func init() {
	rootCmd.AddCommand(harvestCmd)
	harvestCmd.Flags().String("source", sourceAll,
		fmt.Sprintf("Data to harvest: %s, %s, or %s", sourceEvents, sourceYouTubeShorts, sourceAll))
	harvestCmd.Flags().Int("days", youtube.DefaultDays,
		"Number of days in the past to harvest YouTube Shorts (only applies with --source youtubeshorts or all)")
}
