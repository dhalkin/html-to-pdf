# Start from builder
FROM golang:1.17-alpine3.15 AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . .
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/yamato ./yamato/main.go

FROM surnet/alpine-wkhtmltopdf:3.15.0-0.12.6-small
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/yamato --port 8080
