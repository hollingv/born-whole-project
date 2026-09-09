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
	KBURLs      []string // specific pages to include in the knowledge base
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
    KBURLs: []string{
        "https://yourwebsite.org/about",
        "https://yourwebsite.org/research",
    },
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
				EventsUrl:   "https://intactglobal.org/events",
				KBURLs: []string{
					"https://intactglobal.org/about",
					"https://intactglobal.org/initiatives",
					"https://intactglobal.org/litigation",
				},
			},
			{
				Name:        "Intaction",
				Description: "Advancing the health, well-being, and bodily autonomy of boys and men.",
				Website:     "https://intaction.org",
				Thumbnail:   "images/intaction.jpg",
				NewsUrl:     "https://intaction.org/circumcision-news-complete-listing/",
				EventsUrl:   "https://intaction.org/events",
				KBURLs: []string{
					"https://intaction.org",
					"https://intaction.org/ethics-of-circumcision/",
					"https://intaction.org/circumcision-facts/",
				},
			},
			{
				Name:        "Genital Autonomy Legal Defense and Education Fund",
				Description: "To create a world in which the right of everyone to bodily integrity and the freedom to choose what's done to their genitals is legally protected on an equal basis.",
				Website:     "https://www.galdef.org/",
				Thumbnail:   "images/galdef.svg",
				NewsUrl:     "https://www.galdef.org/news",
				EventsUrl:   "https://www.galdef.org/events",
				KBURLs: []string{
					"https://www.galdef.org/about",
					"https://www.galdef.org/legal-defense",
				},
			},
			{
				Name:        "Circumcision Law Reform",
				Description: "Protecting children, youth and parents from the harm of circumcision",
				Website:     "https://circumcisionlawreform.org/",
				Thumbnail:   "images/clr.png",
				NewsUrl:     "https://circumcisionlawreform.org/news",
				EventsUrl:   "https://circumcisionlawreform.org/events",
				KBURLs: []string{
					"https://circumcisionlawreform.org/about",
					"https://circumcisionlawreform.org/legal-reform",
				},
			},
			{
				Name:        "Attorneys For The Rights of the Child",
				Description: "Protecting children, youth and parents from the harm of circumcision",
				Website:     "https://www.arclaw.org/",
				Thumbnail:   "images/arc.jpeg",
				NewsUrl:     "https://www.arclaw.org/news",
				EventsUrl:   "https://www.arclaw.org/events",
				KBURLs: []string{
					"https://www.arclaw.org/about",
					"https://www.arclaw.org/legal-resources",
				},
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
				NewsUrl:     "https://www.doctorsopposingcircumcision.org/news",
				EventsUrl:   "https://www.doctorsopposingcircumcision.org/events",
				KBURLs: []string{
					"https://www.doctorsopposingcircumcision.org/information/",
					"https://www.doctorsopposingcircumcision.org/for-professionals/",
					"https://www.doctorsopposingcircumcision.org/for-parents/",
					"https://www.doctorsopposingcircumcision.org/resources/",
				},
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
				NewsUrl:     "https://www.nocirc.org/news",
				EventsUrl:   "https://www.nocirc.org/events",
				KBURLs: []string{
					"https://www.nocirc.org",
					"https://www.nocirc.org/publish/pamphlet.html",
				},
			},
			{
				Name:        "Intact America",
				Description: "Changing the Way America thinks about circumcision",
				Website:     "https://intactamerica.org",
				Thumbnail:   "images/intactamerica.png",
				NewsUrl:     "https://intactamerica.org/news",
				EventsUrl:   "https://intactamerica.org/events",
				KBURLs: []string{
					"https://intactamerica.org/resources/",
					"https://intactamerica.org/resources/fact-sheets/",
					"https://intactamerica.org/resources/talking-points/",
					"https://intactamerica.org/our-positions/",
				},
			},
			{
				Name:        "Your Whole Baby",
				Description: "The trusted resource for information on circumcision and the foreskin",
				Website:     "https://yourwholebaby.org",
				Thumbnail:   "images/ywb.webp",
				NewsUrl:     "https://yourwholebaby.org/news",
				EventsUrl:   "https://yourwholebaby.org/events",
				KBURLs: []string{
					"https://yourwholebaby.org/researching-parents/",
					"https://yourwholebaby.org/healthcare-providers/",
					"https://yourwholebaby.org/circumcision/",
				},
			},
		},
	},
}
