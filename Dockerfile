# build stage for dealing with go module caching
FROM golang:1.14.3 AS builder
WORKDIR /src

RUN mkdir /build
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN cd cmd/grproxy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -installsuffix netgo -ldflags '-extldflags "-static"'


# final stage
FROM alpine:3.11.6
WORKDIR /app

RUN apk add --no-cache tzdata ca-certificates
COPY --from=builder /build/cmd/grproxy/grproxy /app/grproxy
EXPOSE 9999

ENTRYPOINT ["/app/grproxy"]
