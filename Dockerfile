FROM golang:1.14.2
MAINTAINER Ejmin Vartoumian ( <vvejmin@gmail.com> )

RUN /

EXPOSE $PORT_TO_EXPOSE 1321

RUN go build -o main .

CMD["api.go"]