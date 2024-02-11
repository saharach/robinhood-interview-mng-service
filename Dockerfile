# Start from a minimal Docker image containing Golang runtime
FROM golang:1.22-alpine AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. They will be cached if the go.mod and go.sum files are not changed.
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the binary
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /go/bin/api ./cmd/api/main.go

# Start from a minimal Docker image containing only the runtime
FROM alpine:3.14

# Copy the binary from the build stage
COPY --from=build /go/bin/api /go/bin/api

# Set the current working directory inside the container
WORKDIR /go/bin

# Expose the port that the server will listen on
EXPOSE 8080

# Start the server
CMD ["/go/bin/api"]