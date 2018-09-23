# Build container
FROM golang:1.11-alpine as BUILD

ENV CGO_ENABLED=0

WORKDIR /go/src/app

COPY . .

RUN go install -ldflags="-s -w"

# Release container
FROM scratch as RELEASE

ENV PORT=8080

COPY --from=BUILD /go/bin/app /

EXPOSE 8080
ENTRYPOINT ["/app"]
