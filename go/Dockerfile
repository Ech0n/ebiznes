FROM golang:1.20-alpine

RUN apk update && apk add sqlite gcc musl-dev
ENV CGO_ENABLED=1

WORKDIR /app

COPY ./app /app

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "main.go"]