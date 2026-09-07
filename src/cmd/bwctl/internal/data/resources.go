package data

import (
	"encoding/json"
	"fmt"
	"os"
)

const ResourcesPath = "site/kb/resources.json"

// Resource represents a YouTube Short to feature on the resources page.
type Resource struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoID     string `json:"video_id"`
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
