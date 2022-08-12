FROM golang:1.18.4-alpine3.16

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build ./main.go

CMD [ "./main"]