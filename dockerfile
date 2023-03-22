FROM golang:1.20

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod download

# RUN go run main.go
# RUN go mod download
# RUN go get -u github.com/pressly/goose/cmd/goose
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .
