FROM golang:buster
WORKDIR /app
COPY . .
RUN go mod download
ENV GIN_MODE=release
EXPOSE 8080/tcp
CMD go run main.go