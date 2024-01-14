# docker build --platform linux/amd64 -t buf.build/uneva/protoc-gen-go-http:v1.0.1 .

# Use the official Golang image
FROM golang:1.21-bullseye AS builder

# Set the working directory to the app source code
WORKDIR .

# Ensure we're doing the right thing for Go and set GO111MODULE
ENV GOPATH= CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

# Copy only the go.mod and go.sum files to download dependencies
COPY go.mod .
COPY go.sum .

# Download Go dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o /protoc-gen-go-http .

# Create a minimal image
FROM scratch

# Copy only the binary from the builder stage
COPY --from=builder /protoc-gen-go-http /protoc-gen-go-http

# Set the binary as the entry point
ENTRYPOINT ["/protoc-gen-go-http"]

# Set the user to nobody for better security
USER nobody