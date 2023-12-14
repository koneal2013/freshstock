FROM golang:1.20.5-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /freshstock-api ./cmd

EXPOSE 8080

CMD ["/freshstock-api"]