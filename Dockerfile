# Stage 1: Build
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY app/ ./app/
RUN go build -o server ./app

# Stage 2: Production (distroless)
FROM gcr.io/distroless/static-debian12

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 5000
ENTRYPOINT ["/app/server"]