FROM golang:1.21 as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o demo-app /app/main.go
RUN go mod tidy

FROM alpine

WORKDIR /app
COPY --from=builder /app/demo-app /app/demo-app

EXPOSE 8080

CMD ["/app/demo-app"]

