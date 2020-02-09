FROM golang:buster
WORKDIR /app
COPY . .
RUN go mod download
CMD go run main.go