package data

import (
	"encoding/json"
	"fmt"
	"time"
)

// Date is a calendar date (year, month, day) with no time-of-day component.
// It marshals to/from JSON as a "YYYY-MM-DD" string.
type Date struct {
	time.Time
}

// ParseDate parses a "YYYY-MM-DD" string into a Date.
func ParseDate(s string) (Date, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return Date{}, fmt.Errorf("ParseDate %q: %w", s, err)
	}
	return Date{t}, nil
}

// MustParseDate parses a "YYYY-MM-DD" string and panics on failure.
// Use only in package-level var initializers.
func MustParseDate(s string) Date {
	d, err := ParseDate(s)
	if err != nil {
		panic(err)
	}
	return d
}

// String returns the date as "YYYY-MM-DD", which is what templates render via {{.Date}}.
func (d Date) String() string {
	if d.IsZero() {
		return ""
	}
	return d.Format("2006-01-02")
}

// MarshalJSON encodes the date as a "YYYY-MM-DD" JSON string.
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// flexDateFormats is tried in order by UnmarshalJSON to handle both the
// canonical YYYY-MM-DD storage format and older JSON files that may contain
// raw date strings scraped from source websites (including weekday prefixes).
var flexDateFormats = []string{
	"2006-01-02",
	"Monday, January 2, 2006",
	"Mon, January 2, 2006",
	"Monday, Jan 2, 2006",
	"Mon, Jan 2, 2006",
	"Monday January 2, 2006",
	"Mon Jan 2, 2006",
	"January 2, 2006",
	"Jan 2, 2006",
	"01/02/2006",
	"1/2/2006",
	"January 2006",
}

// UnmarshalJSON decodes a date JSON string into a Date.
// The canonical format is "YYYY-MM-DD"; legacy formats (including those with
// weekday prefixes) are also accepted for backward compatibility.
func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	for _, layout := range flexDateFormats {
		if t, err := time.Parse(layout, s); err == nil {
			d.Time = t
			return nil
		}
	}
	return fmt.Errorf("Date.UnmarshalJSON: unrecognised date format %q", s)
}

// EventKind classifies the origin of an event for display grouping.
type EventKind string

const (
	EventKindFeatured  EventKind = "featured"
	EventKindHarvested EventKind = "harvested"
)

// Label returns a human-readable section heading for the event kind.
func (k EventKind) Label() string {
	switch k {
	case EventKindFeatured:
		return "Featured Events"
	case EventKindHarvested:
		return "Tour & Protest Events"
	default:
		return string(k)
	}
}

// Event represents a single upcoming or past event.
type Event struct {
	Date        Date
	Description string
	URL         string
	OrgID       string
	Kind        EventKind
}

// Org returns the organization associated with this event.
func (e Event) Org() *Organization {
	return FindOrgByID(e.OrgID)
}

// Events is the manually curated list of featured events.
var Events = []Event{
	{
		Date:        MustParseDate("2026-09-14"),
		Description: "Historic Hadachek v. Oregon Court Hearing",
		URL:         "https://www.tickettailor.com/events/intactglobal/2332513",
		OrgID:       "intact-global",
		Kind:        EventKindFeatured,
	},
	{
		Date:        MustParseDate("2027-03-29"),
		Description: "Genital Integrity Awareness Week",
		URL:         "",
		OrgID:       "giaw",
		Kind:        EventKindFeatured,
	},
}
