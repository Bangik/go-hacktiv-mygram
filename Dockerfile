FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o hacktiv-assignment-final

ENTRYPOINT [ "/app/hacktiv-assignment-final" ]

# docker run -it --name hacktiv-ass-final -p 9005:8080 --network laravel-docker-network -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp golang:alpine sh