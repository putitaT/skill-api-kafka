FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN go build -o consumer main.go

FROM alpine
COPY --from=builder /go/src/consumer .

ENTRYPOINT [ "./consumer" ]
