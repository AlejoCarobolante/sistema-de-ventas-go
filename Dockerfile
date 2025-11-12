# Stage 1: Build the Go Application
FROM golang:1.24-alpine AS builder

## Set the working directory inside the container
WORKDIR /app

## Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

## Copy the rest of the application source code
COPY . .

## Build the Go Application
### CGO_ENABLED = 0 ensures a statically linked binary
### -ldflags="-s -w" removes debug information and symbol tables for smaller size
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o app ./cmd

# Stage 2: Create the final lean image
FROM alpine:latest

## Add SSL certificates for HTTPS requests (optional but recommended)
RUN apk --no-cache add ca-certificates

## Set the working directory
WORKDIR /root/

RUN mkdir -p /root/cmd

## Copy the built executable from the builder stage
COPY --from=builder /app/app ./cmd/
COPY --from=builder /app/migrations ./migrations

## Expose the port your application listens on
EXPOSE 8080

## Command to run the application
CMD [ "./cmd/app" ]
