package data

// FAQItem represents a single question and answer.
type FAQItem struct {
	Question string
	Answer   string
}

// FAQItems is the list of frequently asked questions.
var FAQItems = []FAQItem{
	{
		Question: "What is circumcision?",
		Answer:   "Circumcision is the surgical removal of the foreskin, the retractable fold of skin that covers the head of the penis. It is typically performed on infant boys, often within days of birth. While practiced for cultural, religious, and purported medical reasons, it is an irreversible procedure performed on individuals who cannot consent.",
	},
	{
		Question: "What are the medical risks of circumcision?",
		Answer:   "Circumcision carries medical risks including bleeding, infection, scarring, and in rare cases, serious complications requiring further surgery. The foreskin contains thousands of specialized nerve endings and serves protective and sexual functions. Major medical organizations, including the American Academy of Pediatrics, do not recommend routine circumcision, stating that the benefits are not significant enough to justify the procedure.",
	},
	{
		Question: "Does a child have the right to choose?",
		Answer:   "Many advocates and medical ethicists argue that non-therapeutic circumcision of minors violates the child's right to bodily autonomy and informed consent. Since circumcision is irreversible and not medically necessary, they argue the decision should be left to the individual when they are old enough to make an informed choice for themselves. This is the position supported by the organizations featured on this site.",
	},
}
