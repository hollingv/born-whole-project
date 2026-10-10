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
const harvestedDir = "dist/harvested"

const (
	sourceAll           = "all"
	sourceEvents        = "events"
	sourceYouTubeShorts = "youtubeshorts"
)

// writeManifest writes the list of KB text files to generated-kb-file-list.json.
// Paths are stored relative to dist/ (e.g. "kb/intact-global.txt" or "harvested/yt-xxx.txt").
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
	manifestPath := filepath.Join(harvestedDir, "generated-kb-file-list.json")
	if err := os.WriteFile(manifestPath, manifest, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing manifest: %v\n", err)
		return
	}
	fmt.Printf("Written %s\n", manifestPath)
}

var harvestCmd = &cobra.Command{
	Use:   "harvest",
	Short: "Harvest dynamic data and update the knowledge base manifest",
	Long: `Harvests dynamic data and updates dist/harvested/.
Use --source to select which data to harvest: events, youtubeshorts, or all (default).`,
	Run: func(cmd *cobra.Command, args []string) {
		source, _ := cmd.Flags().GetString("source")
		days, _ := cmd.Flags().GetInt("days")

		switch source {
		case sourceAll, sourceEvents, sourceYouTubeShorts:
			// valid
		default:
			fmt.Fprintf(os.Stderr, "Invalid --source %q. Must be one of: %s, %s, %s\n",
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

		writeManifest()
		fmt.Println("Harvest complete.")
	},
}

func init() {
	rootCmd.AddCommand(harvestCmd)
	harvestCmd.Flags().String("source", sourceAll,
		fmt.Sprintf("Data to harvest: %s, %s, or %s", sourceEvents, sourceYouTubeShorts, sourceAll))
	harvestCmd.Flags().Int("days", youtube.DefaultDays,
		"Number of days in the past to harvest YouTube Shorts (only applies with --source youtubeshorts or all)")
}
