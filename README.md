# My Gram Hacktiv

CRUD USER, PHOTO, COMMENT, SOCIAL MEDIA

## Requirements
- Golang
- Postgres

## Links
- [API](http://68.183.183.36:8080)
- [Postman Collection](https://github.com/Bangik/go-hacktiv-mygram/blob/master/Go%20Mygram.postman_collection.json)

## How to run
1. Clone this repository
2. Open terminal and go to the directory
3. Setup environment variable in .env file
4. Run `go mod tidy` to install dependencies
5. Setup database and copy the content of config/database/init.sql
6. Run `go run main.go` to run the program

## How to run with docker-compose
1. Clone this repository
2. Open terminal and go to the directory
3. Setup environment variable in docker-compose.yml file or .env file
4. Run `docker-compose up` to run the program
5. Setup database with
    ```bash
    $ docker exec -it go-db sh
    # psql -U postgres
    # \c mygram
    ```
6. Copy and paste the content of config/database/init.sql to create table