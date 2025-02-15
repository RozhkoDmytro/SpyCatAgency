# Use the official Go image as the build environment
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Install golang-migrate in the builder stage
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Build the application (ensure it is a static binary)
RUN CGO_ENABLED=0 GOOS=linux go build -o spycat-api cmd/api/main.go

# Create the final lightweight image
FROM alpine:latest

# Install necessary dependencies including curl, bash, and postgresql-client
RUN apk --no-cache add ca-certificates bash curl postgresql-client

# Set the working directory
WORKDIR /root/

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/spycat-api .

# Ensure the binary is executable
RUN chmod +x /root/spycat-api

# Install golang-migrate in the final container
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz && mv migrate /usr/local/bin/

# Copy migration files
COPY migrations /migrations

# Copy .env file if it exists
COPY .env .env

# Expose the application port
EXPOSE 8080

# Run migrations before starting the app
CMD ["/usr/local/bin/migrate", "-path", "/migrations", "-database", "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable", "up"] && ["/root/spycat-api"]
