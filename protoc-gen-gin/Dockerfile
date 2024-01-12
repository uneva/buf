# docker build --platform linux/amd64 -t buf.build/vueva/protoc-gen-gin:v1.0.0 .

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
RUN go build -o /protoc-gen-gin .

# Create a minimal image
FROM scratch

# Copy only the binary from the builder stage
COPY --from=builder /protoc-gen-gin /protoc-gen-gin

# Set the binary as the entry point
ENTRYPOINT ["/protoc-gen-gin"]

# Set the user to nobody for better security
USER nobody