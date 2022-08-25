# Start from builder
FROM golang:1.17-alpine3.15 AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . .
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/yamato ./yamato/main.go

FROM surnet/alpine-wkhtmltopdf:3.15.0-0.12.6-small as wkhtmltopdf
FROM openjdk:8-jdk-alpine

# Install dependencies for wkhtmltopdf and libreoffice
RUN apk add --no-cache \
  libstdc++ \
  libx11 \
  libxrender \
  libxext \
  libssl1.1 \
  ca-certificates \
  fontconfig \
  freetype \
  ttf-dejavu \
  ttf-droid \
  ttf-freefont \
  ttf-liberation \
  ttf-ubuntu-font-family \
  libreoffice --no-install-recommends \
&& apk add --no-cache --virtual .build-deps \
  msttcorefonts-installer \
\
# Install microsoft fonts
&& update-ms-fonts \
&& fc-cache -f \
\
# Clean up when done
&& rm -rf /tmp/* \
&& apk del .build-deps

# Copy wkhtmltopdf files from docker-wkhtmltopdf image
COPY --from=wkhtmltopdf /bin/wkhtmltopdf /bin/wkhtmltopdf
COPY --from=wkhtmltopdf /bin/libwkhtmltox* /bin/

COPY --from=build /go/src/app/bin /go/bin
EXPOSE 8080
ENTRYPOINT /go/bin/yamato --port 8080

# only  wkhtmltopdf
#FROM surnet/alpine-wkhtmltopdf:3.15.0-0.12.6-small
#WORKDIR /usr/bin
#COPY --from=build /go/src/app/bin /go/bin
#EXPOSE 8080
#ENTRYPOINT /go/bin/yamato --port 8080