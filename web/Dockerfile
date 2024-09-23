FROM golang:1.19-alpine

# Create a non-root user with a specific UID and GID
RUN addgroup -S appgroup && adduser -S -u 1000 -G appgroup appuser

# Set the working directory
WORKDIR /app

# Copy the application source code
COPY . .

# Download dependencies and build the Go application
RUN go mod tidy
RUN go build -o main .

# Change ownership of the application directory to the non-root user
RUN chown -R appuser:appgroup /app

# Switch to the non-root user
USER 1000

# Expose the application port
EXPOSE 8080

# Set the default command to run the application
CMD ["/app/main"]
