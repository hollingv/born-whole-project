package main

// Organization represents an advocacy organization to feature on the site.
type Organization struct {
	Name        string
	Description string
	Website     string
	Thumbnail   string
	NewsUrl     string
	EventsUrl   string
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
			},
			{
				Name:        "Intaction",
				Description: "Advancing the health, well-being, and bodily autonomy of boys and men.",
				Website:     "https://intaction.org",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://intaction.org/news",
				EventsUrl:   "https://intaction.org/events",
			},
			{
				Name:        "Genital Autonomy Legal Defense and Education Fund",
				Description: "To create a world in which the right of everyone to bodily integrity and the freedom to choose what's done to their genitals is legally protected on an equal basis.",
				Website:     "https://www.galdef.org/",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://www.galdef.org/news",
				EventsUrl:   "https://www.galdef.org/events",
			},
			{
				Name:        "Circumcision Law Reform",
				Description: "Protecting children, youth and parents from the harm of circumcision",
				Website:     "https://circumcisionlawreform.org/",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://circumcisionlawreform.org/news",
				EventsUrl:   "https://circumcisionlawreform.org/events",
			},
			{
				Name:        "Attorneys For The Rights of the Child",
				Description: "Protecting children, youth and parents from the harm of circumcision",
				Website:     "https://www.arclaw.org/",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://www.arclaw.org/news",
				EventsUrl:   "https://www.arclaw.org/events",
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
			},
			{
				Name:        "Intact America",
				Description: "Changing the Way America thinks about circumcision",
				Website:     "https://intactamerica.org",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://intactamerica.org/news",
				EventsUrl:   "https://intactamerica.org/events",
			},
			{
				Name:        "Your Whole Baby",
				Description: "The trusted resource for information on circumcision and the foreskin",
				Website:     "https://yourwholebaby.org",
				Thumbnail:   placeholderImage,
				NewsUrl:     "https://yourwholebaby.org/news",
				EventsUrl:   "https://yourwholebaby.org/events",
			},
		},
	},
}
