FROM golang:1.22 AS build

# Set working directory
WORKDIR /usr/src/app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

# Final stage: Create a lightweight container for running the application
FROM alpine:latest

# Set working directory
WORKDIR /usr/src/app

# Copy the binary executable from the builder stage
COPY --from=build /app /app

EXPOSE 5000

ENTRYPOINT ["/app"]
