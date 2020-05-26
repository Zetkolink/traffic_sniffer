FROM golang

ADD . /go/src/github.com/Zetkolink/traffic_sniffer

RUN go install github.com/Zetkolink/traffic_sniffer

ENTRYPOINT /go/bin/traffic_sniffer