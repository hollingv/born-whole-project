package data

import (
	"encoding/json"
	"fmt"
	"os"
)

const HarvestedEventsPath = "dist/harvested/harvested-events.json"

// LoadHarvestedEvents reads automatically harvested events from disk.
// Returns an empty slice if the file does not exist or cannot be parsed.
func LoadHarvestedEvents() []Event {
	raw, err := os.ReadFile(HarvestedEventsPath)
	if err != nil {
		return nil
	}
	var events []Event
	if err := json.Unmarshal(raw, &events); err != nil {
		fmt.Fprintf(os.Stderr, "[ WARN ] Could not parse %s: %v\n", HarvestedEventsPath, err)
		return nil
	}
	return events
}
