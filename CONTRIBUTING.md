# Contributing to Born Whole Project

Thank you for your interest in contributing. Born Whole Project is an open-source advocacy hub and welcomes contributions from developers, advocates, researchers, and organizations who share our values.

## Ways to contribute

- **Organizations** — if you run an advocacy organization in this space, you can register your site by submitting a pull request to add your details to `src/cmd/bwctl/organizations.go`. 

- **Developers** — the site is built with Go, HTML templates, and Cloudflare Pages. Contributions to the codebase, knowledge base, FAQ, and AI question-answering features are welcome.

- **Researchers and writers** — contributions to the FAQ (`src/cmd/bwctl/faq.go`) and knowledge base sources help improve the accuracy and depth of answers provided to the public.

## Technical overview

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

## Review

All contributions are reviewed before merging to ensure they align with the site's mission. A maintainer will respond to your pull request as soon as possible.
