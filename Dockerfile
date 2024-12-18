# Use original image golang as build stage
FROM golang:1.22-alpine AS build

# Set working directory inside container
WORKDIR /app

# Copy module
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Compile aplikasi Go
RUN go build -o main.out main.go

# Final stage
FROM alpine:3.20.2

WORKDIR /app

RUN apk update && apk --no-cache add ca-certificates

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set working directory inside container

# Copy build binary 

COPY --from=build /app/main.out /app/
COPY --from=build /app/mock/ /app/mock/

# Change ownership of the working directory
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

CMD [ "./main.out" ]
