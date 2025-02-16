FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o wordlesolver cmd/main.go

FROM alpine:latest

COPY --from=builder /app/five_letter_words.txt .
COPY --from=builder /app/wordlesolver .

CMD ["./wordlesolver"]