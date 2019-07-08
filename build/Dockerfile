FROM alpine:3.9 AS base
RUN apk update --no-cache
RUN apk upgrade --no-cache
RUN apk add --no-cache git make curl jq ca-certificates
RUN update-ca-certificates
RUN adduser -D -u 500 -H app
LABEL maintainer="@zephinzer" \
  maintainer_url="https://gitlab.com/usvc" \
  description="see https://gitlab.com/usvc/utils/semver for more information" \
  repository_url="https://gitlab.com/usvc/utils/semver"

FROM golang:alpine AS build
RUN apk update --no-cache
RUN apk upgrade --no-cache
RUN apk add --no-cache make git
WORKDIR /app
COPY . .
RUN make dep
RUN make binaries

FROM base AS semver
USER app
COPY --from=build /app/bin/semver-linux-amd64 /semver
ENTRYPOINT [ "/semver" ]
WORKDIR /repo
LABEL \
  usage="docker run -it -v $(pwd):/repo usvc/semver:latest -h" \
  container_url="https://hub.docker.com/r/usvc/semver" \
  application_url="https://gitlab.com/usvc/utils/semver/tree/master/cmd/semver"

FROM base AS bump
USER app
COPY --from=build /app/bin/bump-linux-amd64 /bump
ENTRYPOINT [ "/bump" ]
WORKDIR /repo
LABEL \
  usage="docker run -it -v $(pwd):/repo usvc/semver-bump:latest -h" \
  container_url="https://hub.docker.com/r/usvc/semver-bump" \
  application_url="https://gitlab.com/usvc/utils/semver/tree/master/cmd/bump"

FROM base AS get
USER app
COPY --from=build /app/bin/get-linux-amd64 /get
ENTRYPOINT [ "/get" ]
WORKDIR /repo
LABEL \
  usage="docker run -it -v $(pwd):/repo usvc/semver-get:latest -h" \
  container_url="https://hub.docker.com/r/usvc/semver-get" \
  application_url="https://gitlab.com/usvc/utils/semver/tree/master/cmd/get"
