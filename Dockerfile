# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY main.go ./

# Build arguments
ARG VERSION=""
ARG COMMIT=""
ARG DATE=""

# Build the application
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags="-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}" -o lancache-adguardhome-sync .

# Final stage
FROM scratch

# Copy SSL certificates for HTTPS requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary from builder stage
COPY --from=builder /app/lancache-adguardhome-sync /lancache-adguardhome-sync

# Set environment variables
ENV ADGUARD_USERNAME="" \
    ADGUARD_PASSWORD="" \
    LANCACHE_SERVER="" \
    ADGUARD_API="" \
    ALL_SERVICES="" \
    SERVICE_NAMES="" \
    SYNC_INTERVAL_MINUTES="1440"

ENTRYPOINT ["/lancache-adguardhome-sync"]