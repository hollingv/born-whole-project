package data

// FAQItem represents a single question and answer.
type FAQItem struct {
	Question string
	Answer   string
}

// FAQItems is the list of frequently asked questions.
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
