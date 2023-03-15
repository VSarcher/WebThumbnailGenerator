from golang:1.20

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy