package data

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

const HarvestedEventsPath = "dist/harvested/harvested-events.json"

// LoadHarvestedEvents reads automatically harvested events from disk.
// Returns an empty slice if the file does not exist or cannot be parsed.
// All loaded events are stamped with EventKindHarvested regardless of what
// the JSON contains, so old files without a Kind field are handled correctly.
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
	for i := range events {
		events[i].Kind = EventKindHarvested
	}
	return events
}

// AllEvents merges the static Events slice with harvested events and returns
// the combined list sorted chronologically by date.
func AllEvents() []Event {
	harvested := LoadHarvestedEvents()
	all := make([]Event, 0, len(Events)+len(harvested))
	all = append(all, Events...)
	all = append(all, harvested...)
	sort.Slice(all, func(i, j int) bool {
		return all[i].Date.Before(all[j].Date.Time)
	})
	return all
}
