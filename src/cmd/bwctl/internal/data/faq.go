package data

// FAQItem represents a single question and answer.
type FAQItem struct {
	Question   string
	Answer     string
	ActionURL  string // optional link shown after the answer
	ActionText string // label for the optional link
}

// FAQItems is the list of frequently asked questions about genital cutting.
var FAQItems = []FAQItem{
	{
		Question: "What is genital cutting?",
		Answer:   "Genital cutting is the surgical removal of the foreskin, the retractable fold of skin that covers the head of the penis. It is typically performed on infant boys, often within days of birth. While practiced for cultural, religious, and purported medical reasons, it is an irreversible procedure performed on individuals who cannot consent.",
	},
	{
		Question: "What are the medical risks of genital cutting?",
		Answer:   "Genital cutting carries medical risks including bleeding, infection, scarring, and in rare cases, serious complications requiring further surgery. The foreskin contains thousands of specialized nerve endings and serves protective and sexual functions. Major medical organizations, including the American Academy of Pediatrics, do not recommend routine genital cutting, stating that the benefits are not significant enough to justify the procedure.",
	},
	{
		Question: "What is circumcision?",
		Answer:   "Circumcision is the surgical removal of the foreskin, a normal and functional part of the male genital area, containing thousands of specialized nerve endings. ",
	},
}

// SiteFAQItems is the list of frequently asked questions about this site.
var SiteFAQItems = []FAQItem{
	{
		Question: "What is Born Whole?",
		Answer:   "Born Whole is an independent, open-source advocacy platform connecting parents, doctors, lawyers, advocates, and the public with the organizations and information working to end forced genital cutting of children.",
	},
	{
		Question: "Who runs this site?",
		Answer:   "Born Whole is an independent project with no affiliation to any single organization. It is maintained by volunteers and open-source contributors who share a commitment to children's bodily autonomy.",
	},
	{
		Question: "How does the AI knowledge base work?",
		Answer:   "The Ask the Knowledge Base feature uses an AI language model that answers questions based exclusively on a curated set of plain-text files sourced from featured advocacy organizations. It does not use the open internet and will tell you if it cannot find relevant information.",
	},
	{
		Question:   "How can I contribute to this site?",
		Answer:     "Born Whole is open source and welcomes contributions from everyone. You can add FAQ entries, submit events, improve organization listings, or contribute code — all via GitHub.",
		ActionURL:  "get-involved.html",
		ActionText: "Visit the Get Involved page",
	},
	{
		Question:   "Is the information on this site medical or legal advice?",
		Answer:     "No. The information on this site is for general educational purposes only and does not constitute medical or legal advice. Always consult a qualified professional for decisions affecting your health or legal situation.",
		ActionURL:  "disclaimers.html",
		ActionText: "Read our Disclaimers",
	},
}
