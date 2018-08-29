FROM golang:1.10

ADD ./ /go/src/jott
WORKDIR /go/src/jott

RUN go get -d -v ./...
RUN go install -v ./...

ENV PORT=3001

CMD ["go", "run", "main.go"]