FROM golang:1.8

ADD ./ /go/src/jott
WORKDIR /go/src/jott

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

ENV PORT=3001

CMD ["./jott"]