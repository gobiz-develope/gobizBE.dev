# Build stage
FROM golang:1.22.6 AS builder
WORKDIR /app

# Copy go mod and go sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Deployment stage
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=builder /app/main /main

# Run the application
CMD ["/main"]
