FROM golang:1.12.9-alpine

RUN apk add git

WORKDIR /go/src/loadbalancer
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["loadbalancer"]