#
# Use the Go image for building the service
#
FROM golang:1.14-alpine AS builder

# Add dependencies and set the directory
RUN apk add --no-progress --no-cache ca-certificates
WORKDIR /go/src/github.com/jedi4z/minesweeper

# Copy the files needed into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service cmd/main.go


#
# Use the alpine image for running the service
#
FROM alpine:latest

# Download dependencies and set the directory
RUN apk --no-cache add ca-certificates
WORKDIR /root/

# Copy outputs from the builder image
COPY --from=builder /go/src/github.com/jedi4z/minesweeper/service ./service

# Run the service application
CMD ["./service"]
