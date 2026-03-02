# First stage: build the Go binary
FROM golang:1.24.11 AS builder

WORKDIR /app

# Copy Go mod and sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app (consider giving the binary a more generic name)
RUN CGO_ENABLED=0 GOOS=linux go build -o go-desent .

# Second stage: minimal runtime
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /app/go-desent /go-desent

# Ensure the binary is executable
RUN chmod +x /go-desent

# Expose port 8082
EXPOSE 8080

# Run the binary
CMD ["/go-desent"]
