## Build
FROM golang:1.19-alpine AS build

ENV GO111MODULE=auto \
    GOPROXY=https://goproxy.cn,direct\
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /

COPY .  .
RUN go mod download

RUN go build -o /PaperTranslation

## Deploy
FROM scratch
WORKDIR /
COPY --from=build /PaperTranslation /PaperTranslation
ENTRYPOINT ["/PaperTranslation"]