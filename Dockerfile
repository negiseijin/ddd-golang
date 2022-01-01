FROM golang:1.17.5-alpine3.15

WORKDIR /go/src/app
COPY ./ /go/src/app/

RUN apk upgrade --update && \
    apk --no-cache add git

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]