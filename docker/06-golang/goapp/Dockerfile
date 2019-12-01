#A hello world example with Go
FROM golang:1.11.5

WORKDIR /go/src/basic-go-app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["basic-go-app"]
