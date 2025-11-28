FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM golang:1.25-alpine

RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port (adjust if your app uses a different port)
EXPOSE 8080

# Run the application
CMD ["./main"]