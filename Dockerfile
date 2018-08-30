FROM golang:1.10

COPY ./ /go/src/github.com/chrisgreg/jott
WORKDIR /go/src/github.com/chrisgreg/jott

RUN go get -v ./

ENV PORT=3001

CMD ["go", "run", "main.go"]