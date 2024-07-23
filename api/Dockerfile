FROM golang:alpine AS builder
WORKDIR /go/src
COPY . .
RUN go build -o api main.go

FROM alpine
COPY --from=builder /go/src/api .

ENTRYPOINT [ "./api" ]
