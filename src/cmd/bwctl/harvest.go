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

// writeManifest writes the list of KB text files to manifest.json.
func writeManifest() {
	files, _ := filepath.Glob(filepath.Join(kbDir, "*.txt"))
	var filenames []string
	for _, f := range files {
		filenames = append(filenames, filepath.Base(f))
	}
	manifest, _ := json.MarshalIndent(filenames, "", "  ")
	manifestPath := filepath.Join(kbDir, "manifest.json")
	if err := os.WriteFile(manifestPath, manifest, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing manifest: %v\n", err)
		return
	}
	fmt.Printf("Written %s\n", manifestPath)
}

var harvestCmd = &cobra.Command{
	Use:   "harvest",
	Short: "Fetch YouTube Shorts and update the knowledge base manifest",
	Long:  `Fetches YouTube Shorts metadata and updates the kb/manifest.json`,
	Run: func(cmd *cobra.Command, args []string) {
		youtube.BuildResources(projectPrefix)
		writeManifest()
		fmt.Println("Harvest complete.")
	},
}

func init() {
	rootCmd.AddCommand(harvestCmd)
}
