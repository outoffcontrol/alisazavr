
FROM golang:1.17 as builder
WORKDIR /go/src/github.com/outoffcontrol/alisazavr
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o alisazavr ./cmd/alisazavr/alisazavr.go

FROM alpine:3.13.5
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/outoffcontrol/alisazavr/alisazavr alisazavr

CMD ["./alisazavr"]