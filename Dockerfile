FROM node:lts-alpine3.13 AS ui_builder
LABEL stage=builder
WORKDIR /app
COPY ui/package*.json ./
RUN npm install
COPY ui/ .
ENV NODE_ENV production
RUN npm run build


FROM golang:1.17.5-alpine AS builder
LABEL stage=builder
WORKDIR /go/src/github.com/wilson-kbs/ks-share

# Add gcc and libc-dev early so it is cached
RUN set -xe \
	&& apk add --no-cache gcc libc-dev

# Install dependencies earlier so they are cached between builds
COPY go.mod go.sum ./
RUN set -xe \
	&& go mod download

# Copy the source code, because directories are special, there are separate layers
COPY cmd/ ./cmd/
COPY --from=ui_builder /app/dist/ ./ui/dist/
COPY ui/ui.go ./ui/ui.go
COPY pkg/ ./pkg/

# Get the version name and git commit as a build argument
ARG GIT_VERSION
ARG GIT_COMMIT

RUN set -xe \
	&& GOOS=linux GOARCH=amd64 go build \
	    -tags prod \
        -o /go/bin/ks-share github.com/wilson-kbs/test-cloudshare-go/cmd/kbs-share

FROM alpine:3.15.0

COPY --from=builder /go/bin/ks-share /usr/local/bin/kbs-share

EXPOSE 8080

ENTRYPOINT ["kbs-share"]
