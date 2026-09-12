# Born Whole Project — build and test toolchain image.
# Contains all tools needed to build, test, and validate the project.
# Built locally via 'make docker-build'. Not published to any registry.

FROM alpine:3.20

ARG LYCHEE_VERSION=0.15.1

RUN apk add --no-cache ca-certificates curl && \
    curl -sSL "https://github.com/lycheeverse/lychee/releases/download/v${LYCHEE_VERSION}/lychee-v${LYCHEE_VERSION}-x86_64-unknown-linux-musl.tar.gz" \
    | tar -xz -C /usr/local/bin lychee

CMD ["lychee", "--version"]
