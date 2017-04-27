FROM ubuntu:14.04 

MAINTAINER liyao.miao@yeepay.com

RUN mkdir -p /go/bin

ADD . /go/bin

WORKDIR /go/bin

ENTRYPOINT ["/go/bin/helloworld"]
