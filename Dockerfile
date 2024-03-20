# Stage 1: Build the Go application
FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hacktiv-assignment-final .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/hacktiv-assignment-final .

COPY --from=builder /app/.env .

ENTRYPOINT [ "./hacktiv-assignment-final" ]


# docker run -it --name hacktiv-ass-final -p 9005:8080 --network laravel-docker-network -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp golang:alpine sh