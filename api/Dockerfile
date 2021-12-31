FROM golang:1.17.5-alpine

WORKDIR /go/src/api

RUN apk update && \
    apk --no-cache add git mysql-client

COPY go.mod ./

RUN go get -u github.com/cosmtrek/air
RUN go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
