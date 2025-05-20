FROM golang:1.24

WORKDIR /server

COPY . .

RUN go mod tidy

EXPOSE 8080 8080

CMD ["go", "run", "main.go"]