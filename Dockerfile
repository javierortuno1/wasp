ARG GOLANG_IMAGE_TAG=1.17-buster

# Build stage
FROM golang:${GOLANG_IMAGE_TAG} AS build

ARG BUILD_TAGS="rocksdb,builtin_static"
ARG BUILD_LD_FLAGS=""

RUN mkdir /wasp
WORKDIR /wasp

# Make sure that modules only get pulled when the module file has changed
COPY go.mod go.sum /wasp/
RUN go mod download
RUN go mod verify

# Project build stage
COPY . .

RUN go build -o . -tags=${BUILD_TAGS} -ldflags="${BUILD_LD_FLAGS}" ./...

# Wasp build
FROM gcr.io/distroless/cc

COPY --from=build /wasp/wasp /usr/bin/wasp