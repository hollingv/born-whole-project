package data

// NewsItem represents a single news article or press mention.
type NewsItem struct {
	Date        string
	Description string
	URL         string
	OrgID       string
}

// Org returns the organization associated with this news item.
func (n NewsItem) Org() *Organization {
	return FindOrgByID(n.OrgID)
}

// NewsItems is the list of news items to display on the news page.
var NewsItems = []NewsItem{
	{
		Date:        "2026-09-15",
		Description: "Oregon September 14 Hearing Update — Intact Global Newsletter Issue 15",
		URL:         "https://www.intactglobal.org/support/newsletters/issue-15-sep-15-2026-oregon-september-14-hearing-update",
		OrgID:       "intact-global",
	},
}
