FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=1 go build -o main .

FROM debian

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/migrations /app/migrations

EXPOSE 8080

CMD ["./main"]