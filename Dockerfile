FROM golang:1.12-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR usr/src/app


COPY . .

# This container exposes port 8080 to the outside world
EXPOSE 8080
