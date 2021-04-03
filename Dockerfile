#build stage
FROM golang:alpine

COPY . /go/src/github.com/z-t-y/tennessgo-web
WORKDIR /go/src/github.com/z-t-y/tennessgo-web

ENV PORT=8080
ENV GOPROXY=https://goproxy.io,direct

RUN go env -w GO111MODULE="on"
RUN go get -u github.com/z-t-y/tennessgo
RUN go get -u github.com/rs/cors
RUN go install

ENTRYPOINT /go/bin/tennessgo-web

EXPOSE 8080
