FROM golang:1.12-alpine

WORKDIR /go/src/github.com/johanhenriksson/d2prox

COPY . .
RUN go build cmd/d2prox.go

CMD ["./d2prox"]