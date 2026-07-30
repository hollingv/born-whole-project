package main

// Organization represents an advocacy organization to feature on the site.
type Organization struct {
	Name        string
	Description string
	Website     string
	Thumbnail   string
}

// OrgGroup represents a category of organizations.
type OrgGroup struct {
	Type          string
	Organizations []Organization
}

const placeholderImage = "images/placeholder.svg"

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
			},
			{
				Name:        "Intaction",
				Description: "Advancing the health, well-being, and bodily autonomy of boys and men.",
				Website:     "https://intaction.org",
				Thumbnail:   placeholderImage,
			},
			{
				Name:        "Intact America",
				Description: "Changing the Way America thinks about circumcision",
				Website:     "https://intactamerica.org",
				Thumbnail:   placeholderImage,
			},
			{
				Name:        "Your Whole Baby",
				Description: "The trusted resource for information on circumcision and the foreskin",
				Website:     "https://yourwholebaby.org",
				Thumbnail:   placeholderImage,
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
			},
			{
				Name:        "Genital Autonomy Legal Defense and Education Fund",
				Description: "To create a world in which the right of everyone to bodily integrity and the freedom to choose what's done to their genitals is legally protected on an equal basis.",
				Website:     "https://www.galdef.org/",
				Thumbnail:   placeholderImage,
			},
			{
				Name:        "NOCIRC",
				Description: "National Organization of Circumcision Information Resource Centers",
				Website:     "https://www.nocirc.org/",
				Thumbnail:   placeholderImage,
			},
			{
				Name:        "Circumcision Law Reform",
				Description: "Protecting children, youth and parents from the harm of circumcision",
				Website:     "https://circumcisionlawreform.org/",
				Thumbnail:   placeholderImage,
			},
			{
				Name:        "Attorneys For The Rights of the Child",
				Description: "Protecting children, youth and parents from the harm of circumcision",
				Website:     "https://circumcisionlawreform.org/",
				Thumbnail:   placeholderImage,
			},
		},
	},
}
