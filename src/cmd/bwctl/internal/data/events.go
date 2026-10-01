package data

// Event represents a single upcoming or past event.
type Event struct {
	Date        string
	Description string
	URL         string
	OrgID       string
}

// Org returns the organization associated with this event.
func (e Event) Org() *Organization {
	return FindOrgByID(e.OrgID)
}

// Events is the list of events to display on the events page.
var Events = []Event{
	{
		Date:        "2026-09-14",
		Description: "Historic Hadachek v. Oregon Court Hearing",
		URL:         "https://www.tickettailor.com/events/intactglobal/2332513",
		OrgID:       "intact-global",
	},
	{
		Date:        "2027-03-29",
		Description: "Genital Integrity Awareness Week",
		URL:         "",
		OrgID:       "giaw",
	},
}
