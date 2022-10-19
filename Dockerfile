# syntax=docker/dockerfile:experimental

# Alpine was chosen for its small footprint comparated to Ubuntu
FROM golang:1.19-alpine

# Define working directory /app
WORKDIR /

# Copy necessary files
COPY go.mod .

# Copy necessary files to /app
COPY . .

# Build the application 
RUN go build -o go-cli

# Run application
CMD /go-cli
