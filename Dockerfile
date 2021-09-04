FROM golang:1.15-alpine  AS builder

RUN apk add --update make

WORKDIR /go/src/github.com/ozonva/ova-entertainment-api/

COPY . /go/src/github.com/ozonva/ova-entertainment-api/

RUN ls -la /go/src/github.com/ozonva/ova-entertainment-api/

RUN make deps && make build

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/ozonva/ova-entertainment-api/bin/ova-entertainment-api .
COPY --from=builder /go/src/github.com/ozonva/ova-entertainment-api/.env .env

RUN chown root:root ova-entertainment-api

EXPOSE 8080
EXPOSE 9090
CMD ["./ova-entertainment-api"]