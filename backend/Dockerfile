# ---- Build stage ----
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# ---- Runtime stage ----
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080
USER nonroot:nonroot

CMD ["./app"]
