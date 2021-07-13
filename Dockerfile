FROM golang:1.15

RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN  go mod download
RUN  go build -o app
ENTRYPOINT ["./app"]
