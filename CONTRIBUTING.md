# Contributing to Born Whole Project

For non-developer contributions — adding organizations, FAQ entries, events, or reporting issues — see the [Get Involved](https://born-whole-project.pages.dev/get-involved.html) page on the site.

## Tech stack

| Component | Technology |
|---|---|
| Site generation | Go (`bwctl site`) |
| Local development server | Go (`go run ./src/server/server.go`) |
| Deployment | Cloudflare Pages |
| AI question answering | Cloudflare Workers AI (`llama-3.1-8b-instruct-fast`) |
| Knowledge base | Plain text files built from organization websites (`bwctl harvest`) |

## Development platform

This project is developed and tested on **Linux**. It will likely also work on **macOS** with little or no modification.

**Windows developers** must use a virtual machine running Ubuntu or a similar Linux distribution or possibly WSL. Nothing has been tested on Windows.

## Getting started

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

## Submitting a contribution

1. Fork this repository on GitHub
2. Create a branch for your change
3. Follow the [Conventional Commits](https://www.conventionalcommits.org/) format for commit messages — this is enforced by a commit-msg hook
4. Submit a pull request with a clear description of your change

## Knowledge base contributions

The knowledge base (`site/kb/`) consists of hand-written plain text files summarising the work and key information of each featured organisation. When contributing to these files, please write original summaries and paraphrases in your own words rather than reproducing verbatim text from third-party websites. Copyright protects the expression of ideas, not the facts or ideas themselves — factual information about circumcision, medical positions, legal cases, and organisational missions may be freely summarised, but copying large passages of text directly from another site is not appropriate.

## Review

All contributions are reviewed before merging to ensure they align with the site's mission. A maintainer will respond to your pull request as soon as possible.
