FROM golang:latest

ADD . /go/src/github.com/viru/weather

WORKDIR /go/src/github.com/viru/weather

RUN go install github.com/viru/weather

EXPOSE 8080
CMD /go/bin/weather