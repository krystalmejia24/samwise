FROM golang:1.13-alpine

WORKDIR /samwise
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD ["go", "run", "cmd/http/main.go"]
