package data

// DonationItem represents a custom donation entry or campaign.
type DonationItem struct {
	Description   string
	URL           string
	OrgName       string
	OrgThumbnail  string
	OrgWebsiteUrl string
}

// DonationItems is the list of custom donation entries to feature on the donate page.
// Organizations with a DonateUrl set in organizations.go are listed automatically.
// Add entries here for specific campaigns, fundraisers, or other custom donation links.
var DonationItems = []DonationItem{
	{
		Description:   "Support Intact Global's legal challenge in Hadachek v. Oregon — a historic case seeking equal protection for all children.",
		URL:           "https://www.intactglobal.org/support/donate",
		OrgName:       "Intact Global",
		OrgThumbnail:  "images/intactglobal.svg",
		OrgWebsiteUrl: "https://intactglobal.org",
	},
}
