package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"born-whole-project/src/cmd/bwctl/internal/data"
	"born-whole-project/src/cmd/bwctl/internal/youtube"
	"born-whole-project/src/cmd/bwctl/knowledge"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

const kbDir = "site/kb"
const maxWorkers = 5

var httpClient = &http.Client{Timeout: 15 * time.Second}

// fetchURL fetches a URL and returns the parsed HTML document.
func fetchURL(u string) (*html.Node, error) {
	resp, err := httpClient.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	return html.Parse(resp.Body)
}

// saveKBFile extracts text from a URL and saves it to the kb directory.
func saveKBFile(orgName, u string) {
	fmt.Printf("Fetching %s (%s)...\n", orgName, u)
	doc, err := fetchURL(u)
	if err != nil {
		fmt.Fprintf(os.Stderr, "  Skipping %s: %v\n", u, err)
		return
	}
	text := knowledge.NormalizeToParagraphs(knowledge.ExtractText(doc))
	outPath := filepath.Join(kbDir, knowledge.URLToFilename(u))
	if err := os.WriteFile(outPath, []byte(text), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "  Error writing %s: %v\n", outPath, err)
		return
	}
	fmt.Printf("  Saved to %s\n", outPath)
}

type kbJob struct {
	orgName string
	url     string
}

// buildTextKB fetches all org KB URLs concurrently using a worker pool.
func buildTextKB() {
	jobs := make(chan kbJob)
	var wg sync.WaitGroup

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				saveKBFile(job.orgName, job.url)
			}
		}()
	}

	for _, group := range data.OrgGroups {
		for _, org := range group.Organizations {
			urls := org.KBURLs
			if len(urls) == 0 {
				urls = []string{org.Website}
			}
			for _, u := range urls {
				jobs <- kbJob{orgName: org.Name, url: u}
			}
		}
	}
	close(jobs)
	wg.Wait()
}

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

const (
	sourceAll           = "all"
	sourceYouTube       = "youtube"
	sourceOrganizations = "organizations"
)

var harvestCmd = &cobra.Command{
	Use:   "harvest",
	Short: "Harvest content from web sources and YouTube into the knowledge base",
	Long:  `Fetches configured web sources and/or YouTube Shorts and saves results to the kb/ directory`,
	Run: func(cmd *cobra.Command, args []string) {
		source, _ := cmd.Flags().GetString("source")

		switch source {
		case sourceAll, sourceYouTube, sourceOrganizations:
			// valid
		default:
			fmt.Fprintf(os.Stderr, "Invalid --source value %q. Must be one of: %s, %s, %s\n", source, sourceAll, sourceYouTube, sourceOrganizations)
			os.Exit(1)
		}

		if err := os.MkdirAll(kbDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating kb directory: %v\n", err)
			os.Exit(1)
		}

		if source == sourceAll {
			// Full rebuild — clear everything
			if err := os.RemoveAll(kbDir); err != nil {
				fmt.Fprintf(os.Stderr, "Error clearing kb directory: %v\n", err)
				os.Exit(1)
			}
			if err := os.MkdirAll(kbDir, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating kb directory: %v\n", err)
				os.Exit(1)
			}
			buildTextKB()
			writeManifest()
		} else if source == sourceOrganizations {
			// Partial rebuild — only remove .txt files, preserve resources.json
			if files, err := filepath.Glob(filepath.Join(kbDir, "*.txt")); err == nil {
				for _, f := range files {
					os.Remove(f)
				}
			}
			buildTextKB()
			writeManifest()
		}

		if source == sourceAll || source == sourceYouTube {
			youtube.BuildResources(projectPrefix)
		}

		fmt.Println("Harvest complete.")
	},
}

func init() {
	rootCmd.AddCommand(harvestCmd)
	harvestCmd.Flags().String("source", sourceAll, fmt.Sprintf("Content source to harvest: %s, %s, %s", sourceAll, sourceYouTube, sourceOrganizations))
}
