// Package data contains the site's content data types and values.
package data

// Organization represents an advocacy organization to feature on the site.
type Organization struct {
	Name        string
	Description string
	Website     string
	Thumbnail   string
	NewsUrl     string
	EventsUrl   string
	DonateUrl   string
}

// OrgGroup represents a category of organizations.
type OrgGroup struct {
	Type          string
	Organizations []Organization
}

const PlaceholderImage = "images/placeholder.jpg"

// ContributeExample is the canonical Go snippet shown on the contribute page.
// Update this whenever the Organization struct fields change.
const ContributeExample = `{
    Name:        "Your Organization Name",
    Description: "A one-sentence description of your work.",
    Website:     "https://yourwebsite.org",
    Thumbnail:   "images/placeholder.jpg",
    NewsUrl:     "https://yourwebsite.org/news",
    EventsUrl:   "https://yourwebsite.org/events",
    DonateUrl:   "https://yourwebsite.org/donate",
},`

// OrgGroups is the list of organization groups to display on the site.
var OrgGroups = []OrgGroup{
	{
		Type: "Legal",
		Organizations: []Organization{
			{
				Name:        "Intact Global",
				Description: "Support our mission to protect all children from forced, non-religious genital cutting.",
				Website:     "https://intactglobal.org",
				Thumbnail:   "images/intactglobal.svg",
				NewsUrl:     "https://intactglobal.org/press",
				EventsUrl:   "",
				DonateUrl:   "https://www.intactglobal.org/support/donate",
			},
			{
				Name:        "Intaction",
				Description: "Advancing the health, well-being, and bodily autonomy of boys and men.",
				Website:     "https://intaction.org",
				Thumbnail:   "images/intaction.jpg",
				NewsUrl:     "https://intaction.org/circumcision-news-complete-listing/",
				EventsUrl:   "",
				DonateUrl:   "https://www.zeffy.com/en-US/embed/donation-form/donate-to-make-a-difference-3786?modal=true",
			},
			{
				Name:        "Genital Autonomy Legal Defense and Education Fund",
				Description: "To create a world in which the right of everyone to bodily integrity and the freedom to choose what's done to their genitals is legally protected on an equal basis.",
				Website:     "https://www.galdef.org/",
				Thumbnail:   "images/galdef.svg",
				NewsUrl:     "https://www.galdef.org/news/",
				EventsUrl:   "",
				DonateUrl:   "https://www.galdef.org/donate/",
			},
			{
				Name:        "Circumcision Law Reform",
				Description: "Protecting children, youth and parents from the harm of circumcision",
				Website:     "https://circumcisionlawreform.org/",
				Thumbnail:   "images/clr.png",
				NewsUrl:     "",
				EventsUrl:   "",
				DonateUrl:   "",
			},
			{
				Name:        "Attorneys For The Rights of the Child",
				Description: "Protecting children, youth and parents from the harm of circumcision",
				Website:     "https://www.arclaw.org/",
				Thumbnail:   "images/arc.jpeg",
				NewsUrl:     "https://www.arclaw.org/news",
				EventsUrl:   "https://www.arclaw.org/events",
				DonateUrl:   "https://www.arclaw.org/donate",
			},
		},
	},
	{
		Type: "Medical",
		Organizations: []Organization{
			{
				Name:        "Doctors Opposing Circumcision",
				Description: "An international network of physicians dedicated to protecting the genital integrity and eventual autonomy of all children",
				Website:     "https://www.doctorsopposingcircumcision.org/",
				Thumbnail:   "images/doc.svg",
				NewsUrl:     "",
				EventsUrl:   "",
				DonateUrl:   "",
			},
		},
	},
	{
		Type: "Informational",
		Organizations: []Organization{
			{
				Name:        "NOCIRC",
				Description: "National Organization of Circumcision Information Resource Centers",
				Website:     "https://www.nocirc.org/",
				Thumbnail:   PlaceholderImage,
				NewsUrl:     "",
				EventsUrl:   "",
				DonateUrl:   "",
			},
			{
				Name:        "Intact America",
				Description: "Changing the Way America thinks about circumcision",
				Website:     "https://intactamerica.org",
				Thumbnail:   "images/intactamerica.png",
				NewsUrl:     "",
				EventsUrl:   "https://intactamerica.org/events",
				DonateUrl:   "https://intactamerica.org/donate/",
			},
			{
				Name:        "Your Whole Baby",
				Description: "The trusted resource for information on circumcision and the foreskin",
				Website:     "https://yourwholebaby.org",
				Thumbnail:   "images/ywb.webp",
				NewsUrl:     "",
				EventsUrl:   "",
				DonateUrl:   "https://www.yourwholebaby.org/donate",
			},
		},
	},
	{
		Type: "Public Outreach",
		Organizations: []Organization{
			{
				Name:        "Prevail Over The System",
				Description: "Connecting Intactivism with businesses",
				Website:     "https://www.prevailoverthesystem.com",
				Thumbnail:   PlaceholderImage,
				NewsUrl:     "",
				EventsUrl:   "",
				DonateUrl:   "https://ko-fi.com/potsltd",
			},
			{
				Name:        "Blood Stained Men",
				Description: "To warn the American people that circumcision is cruel, worthless, and destructive",
				Website:     "https://bloodstainedmen.com",
				Thumbnail:   "images/bsm.png",
				NewsUrl:     "",
				EventsUrl:   "https://www.bloodstainedmen.com/about-us/events/",
				DonateUrl:   "https://www.bloodstainedmen.com/donate/",
			},
		},
	},
}
