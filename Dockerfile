# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine AS build

ENV GO111MODULE=auto \
    GOPROXY=https://goproxy.cn,direct\
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .
RUN go mod download

COPY *.go ./

RUN go build -o /PaperTranslation

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /PaperTranslation /PaperTranslation

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/PaperTranslation"]