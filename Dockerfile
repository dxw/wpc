FROM golang:1.13

WORKDIR /go/src/github.com/wpc
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

WORKDIR /app
ENTRYPOINT ["wpc"]
