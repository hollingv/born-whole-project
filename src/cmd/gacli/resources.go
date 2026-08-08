package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const resourcesPath = "site/kb/resources.json"
const youtubeChannel = "AttorneyClopper"

// Resource represents a YouTube Short to feature on the resources page.
type Resource struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoID     string `json:"video_id"`
}

// loadResources reads resources from content/resources.json.
// Returns an empty slice if the file does not exist or cannot be parsed.
func loadResources() []Resource {
	data, err := os.ReadFile(resourcesPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ WARN ] Could not read %s: %v\n", resourcesPath, err)
		return nil
	}
	var resources []Resource
	if err := json.Unmarshal(data, &resources); err != nil {
		fmt.Fprintf(os.Stderr, "[ WARN ] Could not parse %s: %v\n", resourcesPath, err)
		return nil
	}
	return resources
}
