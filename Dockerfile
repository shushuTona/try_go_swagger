FROM quay.io/goswagger/swagger AS swagger

FROM golang:1.21

COPY --from=swagger /usr/bin/swagger /usr/bin

WORKDIR /go/src
