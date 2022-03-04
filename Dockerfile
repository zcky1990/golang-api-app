FROM golang:1.12-alpine

RUN apk add --no-cache git

WORKDIR /app/go-api-app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o ./dist/go-api-app .

EXPOSE 10000

CMD ["./dist/go-api-app"]