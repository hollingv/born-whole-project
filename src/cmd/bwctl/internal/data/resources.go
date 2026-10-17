package data

import (
	"encoding/json"
	"fmt"
	"os"
)

// ResourcesPath is the location of the JSON file written by bwctl harvest containing harvested YouTube Shorts metadata.
const ResourcesPath = "dist/harvested/harvested-youtube-shorts.json"

// Resource represents a YouTube Short to feature on the resources page.
type Resource struct {
	Title       string `json:"title"`       // video title as returned by the YouTube Data API
	Description string `json:"description"` // short video description
	VideoID     string `json:"video_id"`    // YouTube video ID used to construct embed and watch URLs
}

// LoadResources reads resources from site/kb/resources.json.
// Returns an empty slice if the file does not exist or cannot be parsed.
func LoadResources() []Resource {
	raw, err := os.ReadFile(ResourcesPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ WARN ] Could not read %s: %v\n", ResourcesPath, err)
		return nil
	}
	var resources []Resource
	if err := json.Unmarshal(raw, &resources); err != nil {
		fmt.Fprintf(os.Stderr, "[ WARN ] Could not parse %s: %v\n", ResourcesPath, err)
		return nil
	}
	return resources
}
