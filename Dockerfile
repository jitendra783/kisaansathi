# ============================================
# Stage 1: Build
# ============================================
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Required packages
RUN apk add --no-cache git ca-certificates

# Copy dependency files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/kisansathi-backend ./cmd/main.go


# ============================================
# Stage 2: Runtime
# ============================================
FROM alpine:3.22

WORKDIR /app

# CA certificates for HTTPS/API calls
RUN apk add --no-cache ca-certificates

# Copy binary
COPY --from=builder /app/kisansathi-backend .

# Application port
EXPOSE 8008

# Start application
CMD ["./kisansathi-backend"]