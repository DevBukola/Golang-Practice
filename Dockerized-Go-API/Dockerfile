FROM golang:tip-alpine3.22
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["go", "run", "main.go"]