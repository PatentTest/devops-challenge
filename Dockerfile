FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY app/go.mod app/go.sum ./
RUN go mod download

COPY app/ .
RUN go build -o server .

# Stage 2: Production (distroless)
FROM gcr.io/distroless/static-debian12

WORKDIR /app
COPY --from=builder /app/server ./

EXPOSE 5000
ENTRYPOINT ["/app/server"]