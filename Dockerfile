# Use an official Go image as the base image
FROM golang:latest

ENV PORT=8080
ENV DB_MYSQL_CONNECTION=root:root@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local
ENV JWT_SECRET_KEY=WE5CcFRaTlpON3dXZ2szYVV5eTM=

# Set the working directory in the container to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build -o dist ./src/cmd/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./dist"]
