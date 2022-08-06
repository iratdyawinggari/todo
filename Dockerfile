FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /todo

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]