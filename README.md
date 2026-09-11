# Born Whole Project

**Every child deserves the right to bodily autonomy.**

Born Whole Project is an open-source advocacy hub dedicated to raising awareness about the harms of genital cutting, connecting people with trusted organizations, and empowering parents and the general public to make informed decisions.

---
## Our Mission

We believe that access to accurate, compassionate information can change outcomes for children around the world. Born Whole Project serves as a central resource for:

- **Education** — helping parents, healthcare providers, and the general public understand the medical, ethical, and legal dimensions of genital cutting
- **Connection** — directing people to the organizations leading the fight to protect children's bodily integrity
- **Action** — supporting the legal and advocacy efforts that are working to end this practice

---
## Collaborative Development

This project is built openly on GitHub and welcomes contributions from developers, advocates, researchers, and organizations who share our values.

See [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to get involved.

---
## Sharing News and Events

A core goal of Born Whole Project is to become the go-to hub for news and events in the intactivist space. Each organization featured on the site is encouraged to publish updates at consistent URLs:

| Path | Content |
|---|---|
| `/news` | Recent news and press coverage |
| `/events` | Upcoming events, conferences, and actions |
| `/research` | Academic and medical research |
| `/legal-cases` | Active and completed legal cases |

The site aggregates this content automatically — organizations that adopt these URL patterns are integrated into the hub without any manual effort.

---
## Educating Parents and the Public

The harm of genital cutting is not widely understood. Many parents make decisions under social pressure, misinformation, or without realizing there is a choice to be made. Born Whole Project addresses this through:

- **Plain-language explanations** of what circumcision is, what the medical evidence says, and what the ethical arguments against it are.
- **An AI-powered Ask feature** that answers natural language questions using content sourced directly from vetted advocacy organizations — grounded in real information, not assumptions
- **A FAQ page** covering the most common questions parents ask
- **Direct links** to the organizations best placed to provide support, guidance, and legal resources

---
## Technical Overview

| Component | Technology |
|---|---|
| Site generation | Go (`bwctl site`) |
| Local development server | Go (`go run ./src/server/server.go`) |
| Deployment | Cloudflare Pages |
| AI question answering | Cloudflare Workers AI (`llama-3.1-8b-instruct-fast`) |
| Knowledge base | Plain text files built from organization websites (`bwctl harvest`) |

### Development platform

This project is developed and tested on **Linux**. It will likely also work on **macOS** with little or no modification.

**Windows developers** must use a virtual machine running Ubuntu or a similar Linux distribution or possibly WSL.  Nothing has been tested on Windows. 

### Getting started

```sh
# Install dependencies
make init

# Build the site and run locally
make build
go run ./src/server/server.go

# Harvest content from organization web sources and YouTube
./bwctl harvest

# Run tests
make test
```

---

## Learn. Understand. Take action.

If you are a parent researching this topic, a healthcare provider looking for resources, or an advocate wanting to get involved — you are in the right place.

Visit the organizations listed on this site. Read the FAQ. Ask a question. And share what you learn.
---

*Born Whole Project is an independent advocacy resource and is not affiliated with any single organization. We feature organizations whose work we believe deserves greater visibility.*
