FROM golang:1.18.4-alpine3.16

ARG WORK_DIR

WORKDIR $WORK_DIR

COPY . .

RUN apk add --no-cache ca-certificates \
    bash \
    sudo \
    git 

RUN go install github.com/cosmtrek/air@v1.40.4
RUN go install github.com/mgechev/revive@v1.2.1

ENTRYPOINT [ "./entrypoint.sh" ]