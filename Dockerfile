FROM golang:1.13-alpine3.10 AS builder

WORKDIR /go/src/gihub.com/Zetkolink/traffic_sniffer
COPY . .

RUN apk add --update gcc libc-dev libpcap-dev

RUN go install

FROM alpine:3.10

RUN apk add --no-cache libc-dev libpcap-dev

COPY --from=builder /go/bin /usr/bin/

ENV DEVICE="lo"

CMD [ "/usr/bin/traffic_sniffer" ]