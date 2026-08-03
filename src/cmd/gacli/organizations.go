package main

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

const placeholderImage = "images/placeholder.jpg"

// orgGroups is the list of organization groups to display on the site.
var orgGroups = []OrgGroup{
	{
		Type: "Legal",
		Organizations: []Organization{
			{
				Name:        "Intact Global",
				Description: "Leading the legal efforts to end the practice of circumcision by protecting all children from genital cutting",
				Website:     "https://intactglobal.org",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://intactglobal.org/news",
				EventsUrl:   "https://intactglobal.org/events",
				KBURLs: []string{
					"https://intactglobal.org/research",
					"https://intactglobal.org/legal-cases",
					"https://intactglobal.org/about",
				},
			},
			{
				Name:        "Intaction",
				Description: "Advancing the health, well-being, and bodily autonomy of boys and men.",
				Website:     "https://intaction.org",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://intaction.org/news",
				EventsUrl:   "https://intaction.org/events",
				KBURLs: []string{
					"https://intaction.org/research",
					"https://intaction.org/legal-cases",
					"https://intaction.org/about",
				},
			},
			{
				Name:        "Genital Autonomy Legal Defense and Education Fund",
				Description: "To create a world in which the right of everyone to bodily integrity and the freedom to choose what's done to their genitals is legally protected on an equal basis.",
				Website:     "https://www.galdef.org/",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://www.galdef.org/news",
				EventsUrl:   "https://www.galdef.org/events",
				KBURLs: []string{
					"https://www.galdef.org/research",
					"https://www.galdef.org/legal-cases",
					"https://www.galdef.org/about",
				},
			},
			{
				Name:        "Circumcision Law Reform",
				Description: "Protecting children, youth and parents from the harm of circumcision",
				Website:     "https://circumcisionlawreform.org/",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://circumcisionlawreform.org/news",
				EventsUrl:   "https://circumcisionlawreform.org/events",
				KBURLs: []string{
					"https://circumcisionlawreform.org/research",
					"https://circumcisionlawreform.org/legal-cases",
					"https://circumcisionlawreform.org/about",
				},
			},
			{
				Name:        "Attorneys For The Rights of the Child",
				Description: "Protecting children, youth and parents from the harm of circumcision",
				Website:     "https://www.arclaw.org/",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://www.arclaw.org/news",
				EventsUrl:   "https://www.arclaw.org/events",
				KBURLs: []string{
					"https://www.arclaw.org/research",
					"https://www.arclaw.org/legal-cases",
					"https://www.arclaw.org/about",
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
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://www.doctorsopposingcircumcision.org/news",
				EventsUrl:   "https://www.doctorsopposingcircumcision.org/events",
				KBURLs: []string{
					"https://www.doctorsopposingcircumcision.org/research",
					"https://www.doctorsopposingcircumcision.org/legal-cases",
					"https://www.doctorsopposingcircumcision.org/about",
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
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://www.nocirc.org/news",
				EventsUrl:   "https://www.nocirc.org/events",
				KBURLs: []string{
					"https://www.nocirc.org/research",
					"https://www.nocirc.org/legal-cases",
					"https://www.nocirc.org/about",
				},
			},
			{
				Name:        "Intact America",
				Description: "Changing the Way America thinks about circumcision",
				Website:     "https://intactamerica.org",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://intactamerica.org/news",
				EventsUrl:   "https://intactamerica.org/events",
				KBURLs: []string{
					"https://intactamerica.org/research",
					"https://intactamerica.org/legal-cases",
					"https://intactamerica.org/about",
				},
			},
			{
				Name:        "Your Whole Baby",
				Description: "The trusted resource for information on circumcision and the foreskin",
				Website:     "https://yourwholebaby.org",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://yourwholebaby.org/news",
				EventsUrl:   "https://yourwholebaby.org/events",
				KBURLs: []string{
					"https://yourwholebaby.org/research",
					"https://yourwholebaby.org/legal-cases",
					"https://yourwholebaby.org/about",
				},
			},
		},
	},
}
