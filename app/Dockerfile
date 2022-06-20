# Start from
FROM golang:1.17-alpine3.15 AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY .. .
#RUN go mod init github.com/dhalkin/html-to-pdf
#RUN go mod tidy
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/yamato ./cmd/yamato/main.go

FROM alpine:3.15
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/yamato --port 8080
