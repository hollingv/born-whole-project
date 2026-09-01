# Bodily Integrity Commons: Shared Infrastructure for the Intactivist Movement

**Document type:** White Paper  
**Status:** Active Development  
**Deployment:** Cloudflare Pages  

---

## Abstract

This paper describes Bodily Integrity Commons (BIC), an open-source web platform serving as shared digital infrastructure for the intactivist movement. BIC aggregates information from advocacy organizations, provides AI-assisted public education, and establishes a collaborative protocol through which organizations can contribute to and benefit from a shared resource hub. Built on GitHub, deployed on Cloudflare, and developed entirely in the open, BIC demonstrates how distributed collaboration can create durable, trustworthy infrastructure for the bodily integrity movement.  This paper describes existing functionality as well as future plans for the project.

---

## 1. Introduction

The movement to end the genital cutting of minors is global, growing, and fragmented. Dozens of organizations operate independently — each maintaining their own website, publishing their own research, news and events, and reaching their own audience. This fragmentation limits impact: a parent seeking information may not know where to look, a researcher may not find the most relevant evidence, and an advocate may not know which organizations are most active in their region.

BIC addresses this fragmentation *not by replacing existing organizations, but by providing a shared platform that aggregates, connects, and amplifies their work.* BIC is conceived as infrastructure: like a road network or an internet protocol, its value grows with the number of participants.

---

## 2. The Problem: Fragmented Information, Limited Reach

Parents making decisions about infant circumcision frequently report difficulty finding reliable, accurate information. Medical and legal positions on the practice vary widely across countries and institutions, and advocacy organizations — though numerous and well-informed — often have limited web visibility.

Several structural problems compound this:

- **Discovery:** No central directory connects the public with the organizations working in this space.
- **Duplication:** Organizations independently produce similar educational content without a mechanism to share or cross-reference it.
- **Accessibility:** Technical content (medical research, legal arguments) is not easily accessible to general audiences.
- **Sustainability:** Individual organizations face resource constraints that limit their technical capacity.

BIC addresses each of these problems through a collaborative, open-source model.

---

## 3. The Open Source Model

BIC is developed entirely in the open on GitHub under an open-source license. This is not incidental — it is the foundation of the project's credibility, resilience, and collaborative potential.

### 3.1 Transparency and Trust

Any person can inspect the complete source code and knowledge base, audit the AI prompts used to answer questions, and confirm that the knowledge base is drawn from vetted sources. In a domain where trust is paramount, open code is a meaningful signal of accountability.

### 3.2 Resilience

An open-source project hosted on a distributed platform cannot be silenced by any single actor. The repository can be forked, the site can be redeployed, and the codebase can be continued by any qualified contributor. This resilience is particularly important for advocacy projects that may face opposition or lose individual maintainers.

### 3.3 Collaborative Development

GitHub's pull request model provides a structured mechanism for all contributions. Organizations can add themselves to the platform with a simple file edit and pull request — no technical infrastructure of their own is required. The review process ensures quality while keeping participation accessible.

> "The pull request model means every addition is reviewed, every organization is vetted, and the history of every change is permanently recorded on a public ledger."

---

## 4. The Dispatching Hub Architecture

BIC operates as a dispatching hub — a central resource that directs users to the most relevant organizations and content, rather than attempting to replicate or compete with existing work.

### 4.1 The URL Convention Protocol

BIC defines a lightweight protocol: a standard set of URL paths that participating organizations are encouraged to publish content under. Organizations that adopt these conventions are automatically integrated into BIC's aggregation and AI features.

| Path | Expected Content | BIC Feature |
|---|---|---|
| `/about` | Organization background | Knowledge base |
| `/research` | Academic & medical research | Knowledge base, AI answers |
| `/legal-cases` | Active and completed cases | Legal tracker (planned) |
| `/news` | Press coverage | News aggregator (planned) |
| `/events` | Upcoming events | Events calendar (planned) |

This model is analogous to RSS — organizations that publish a standard feed are automatically included in feed readers. BIC functions as the feed reader for the intactivist movement.

### 4.2 The Knowledge Base

The `bictl harvest` command fetches content from each registered organization's designated pages, extracts meaningful prose, and stores it as committed text files. This knowledge base powers the AI question-answering feature and can be updated at any time by running the harvest command and committing the results.

---

## 5. Technical Architecture

BIC is designed for low operational cost, high reliability, and ease of contribution. The architecture favours static generation over dynamic server-side rendering, serverless functions over managed servers, and committed content in GitHub over live database queries.

| Component | Technology | Role |
|---|---|---|
| Site generation | Go / `bictl site` | Renders HTML from templates and data at build time |
| Local development | Go (`src/server`) | Full server with AI and KB search for local testing |
| Hosting & CDN | Cloudflare Pages | Global distribution, zero-ops deployment |
| AI answering | Cloudflare Workers AI | LLM-powered answers grounded in the KB |
| Content harvest | Go / `bictl harvest` | Fetches org content and YouTube Shorts |
| CI/CD | GitHub Actions | Test, build, deploy and integration-test on push |
| Versioning | Conventional Commits + cog | Automated changelog and semantic versioning |
| Automated Testing | Go testing framework | Automated testing on any push |

### 5.1 AI Question Answering

BIC implements a Retrieval-Augmented Generation (RAG) pipeline. When a user submits a question, meaningful keywords are extracted and the most relevant paragraphs from the knowledge base are identified. These paragraphs, together with the question, are sent to Cloudflare Workers AI, which generates a natural language response *grounded exclusively in the curated knowledge base.* Sources are cited in every response, and a disclaimer reminds users to verify information with the linked organizations.

### 5.2 Deployment and Quality Assurance

Every push to the `main` branch triggers an automated pipeline: unit tests, integration tests, site generation, deployment to Cloudflare Pages, and post-deployment integration tests against the live URL. A parallel pipeline serves the `devel` branch at a separate preview URL which serves as a testing ground before merging to `main`. Conventional commit messages are enforced by a `commit-msg` git hook, and all commits are validated before release using the `cog` tool.

---

## 6. Participation Model

BIC is structured to accommodate contributors at every level of technical proficiency that range from writers, UI designers, to developers.

### 6.1 Writers and Organizations

Any advocacy organization can add itself to BIC by editing a single source file on GitHub and submitting a pull request. The Contribute page on the BIC website provides step-by-step instructions requiring no local development environment. Organizations that additionally adopt BIC's URL conventions receive automatic integration into the knowledge base and future aggregation features.

### 6.2 UI Designers

UI designers may contribute to the design of the BIC website and UI components. They are responsible for creating visually appealing and user-friendly interfaces.

### 6.3 Developers

Developers may contribute to any layer of the stack — Go backend, HTML templates, CSS, JavaScript, CI/CD configuration, or AI integration. The project uses standard Go tooling, a GNU Makefile for common tasks, and GitHub Actions for automation. A `make test` command runs the full test suite locally.

### 6.4 Advocates and Researchers

Non-technical contributors play an equally important role. Identifying high-quality pages to include in the knowledge base, curating YouTube educational content, contributing FAQ entries, and verifying organization information are all valuable contributions that require no programming knowledge.

---

## 7. Strategic Vision

BIC's long-term vision is to become the canonical digital infrastructure for the intactivist movement — a platform that every organization, advocate, researcher, and parent can rely on, and that grows stronger with each new participant.

### 7.1 Near-term Development

- Semantic search using vector embeddings (Cloudflare Vectorize) to improve the accuracy of AI-generated answers
- Paragraph-level chunking and improved knowledge base quality to provide richer AI context
- Expanded YouTube resources, filtered and curated by topic
- Streaming AI responses for improved user experience

### 7.2 Medium-term Development

- News aggregation page drawing from registered organizations' `/news` paths
- Events calendar aggregating from organizations' `/events` paths
- Legal case tracker drawing from `/legal-cases` pages
- Conformance indicators showing which organizations have adopted the URL protocol

### 7.3 The Network Effect

The value of BIC's infrastructure model compounds as participation grows. Each additional organization enriches the knowledge base, expands the directory, and increases the platform's utility for the public. Each developer contribution improves the experience for all users. This network effect is the core strategic rationale for the open-source, collaborative approach.

> "BIC is not owned by any single organization. It is shared infrastructure, built by the movement, for the movement. Every contribution — however small — makes it stronger."

---

## 8. Conclusion

Bodily Integrity Commons represents a new model for advocacy infrastructure: open-source, collaboratively maintained, AI-augmented, and designed to grow with the movement it serves. By providing shared technical infrastructure, BIC enables individual organizations to focus on their core work while collectively achieving greater reach, credibility, and impact than any could achieve independently.

The platform is operational, deployed globally, and open to contributions today. Organizations wishing to participate may submit a pull request to the public GitHub repository or visit the Contribute page at the BIC website for guided instructions.

The code is open. The movement is welcome.

---

## 9. References and Resources (tbd)

- **Repository:** 
- **Website:** 
- **Contribute:** 
- **Organizations:** 
- **Cloudflare Pages:** 
- **Cloudflare Workers AI:**
