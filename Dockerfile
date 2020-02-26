FROM golang:buster
WORKDIR /app
COPY . .
RUN go mod download
ENV GIN_MODE=release
CMD go run main.go