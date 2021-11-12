FROM golang:latest

LABEL Maintainer="JOSE AVILA <jose.avila@globant.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .
