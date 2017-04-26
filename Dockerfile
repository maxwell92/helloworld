FROM img.reg.3g:15000/ubuntu-base:v3

MAINTAINER liyao.miao@yeepay.com

RUN mkdir -p /go/bin

ADD . /go/bin

WORKDIR /go/bin

ENTRYPOINT ["/go/bin/helloworld"]
