FROM golang:1.17.3

RUN mkdir /go/src/app
COPY . /go/src/app
WORKDIR /go/src/app

RUN go build -o /main

EXPOSE 8080

CMD ["/main"]
