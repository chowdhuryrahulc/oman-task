# Stage 1: Build the binary
FROM golang:1.20.3-alpine3.17 AS builder
WORKDIR /app   
COPY . .
RUN go build -o oman-task .

# Stage 2: Create the final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app .
EXPOSE 4000
ENTRYPOINT ["/app/oman-task"]