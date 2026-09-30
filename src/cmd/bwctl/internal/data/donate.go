package data

// DonationItem represents a custom donation entry or campaign.
type DonationItem struct {
	Description string
	URL         string
	OrgID       string
}

// Org returns the organization associated with this donation item.
func (d DonationItem) Org() *Organization {
	return FindOrgByID(d.OrgID)
}

// DonationItems is the list of custom donation entries to feature on the donate page.
// Organizations with a DonateUrl set in organizations.go are listed automatically.
// Add entries here for specific campaigns, fundraisers, or other custom donation links.
var DonationItems = []DonationItem{
	{
		Description: "Support Intact Global's legal challenge in Hadachek v. Oregon — a historic case seeking equal protection for all children.",
		URL:         "https://www.intactglobal.org/support/donate",
		OrgID:       "intact-global",
	},
	{
		Description: "Back Intact Global's Fight for Equal Protection in Colorado",
		URL:         "https://www.gofundme.com/f/back-intact-globals-fight-for-equal-protection-in-colorado",
		OrgID:       "intact-global",
	},
	{
		Description: "Critical Call to Support Boys' and Men's Well Being",
		URL:         "https://www.zeffy.com/en-US/donation-form/hec-washington-dc?modal=true",
		OrgID:       "intaction",
	},
}
