FROM golang:1.19-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN  go build -o build/coffee-rest main.go

EXPOSE 50051

CMD [ "./build/coffee-rest" ]