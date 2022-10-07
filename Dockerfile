FROM golang:1.19.1-alpine as ci
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.49.0

FROM golang:1.19.1-alpine as builder
WORKDIR /go/src/github.com/GorginZ/community-library
COPY . /go/src/github.com/GorginZ/community-library

RUN go mod download
RUN go mod tidy
ENV CGO_ENABLED=0

ARG version

RUN CGO_ENABLED=0 go build -v -o /_build/community-library -ldflags "-X 'github.com/GorginZ/community-library/metadata.Version=${version}'"


FROM scratch
COPY --from=builder /_build/community-library /app
ENTRYPOINT ["/app"]
