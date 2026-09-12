package data

// Event represents a single upcoming or past event.
type Event struct {
	Date        string
	Description string
	URL         string
}

// Events is the list of events to display on the events page.
var Events = []Event{
	{
		Date:        "2026-09-14",
		Description: "Historic Hadachek v. Oregon Court Hearing",
		URL:         "https://www.tickettailor.com/events/intactglobal/2332513#",
	},
}
