FROM golang:1.13.5-alpine3.10
MAINTAINER srsh @raben2 rabenstein@srsh.io

ENV EMBY_URL = $EMBY_URL
ENV EMBY_USER = $EMBY_USER
ENV EMBY_TOKEN = EMBY_TOKEN

RUN mkdir -p /go/src/github.com/raben2/embyoracle
COPY . /go/src/github.com/raben2/embyoracle
WORKDIR /go/src/github.com/raben2/embyoracle
RUN go get .../.
RUN go build -o /bin/embyoracle 

WORKDIR /

ENTRYPOINT ["/bin/embyoracle"]