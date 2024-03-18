FROM golang:alpine

WORKDIR /app

RUN go mod download

COPY . .

# docker run -it --name hacktiv-ass-final -p 9005:8080 --network laravel-docker-network -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp golang:alpine sh