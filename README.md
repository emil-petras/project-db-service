# db-service

## list of related repositories
https://github.com/emil-petras/project-web-service
https://github.com/emil-petras/project-idempotency-service
https://github.com/emil-petras/project-proto

this service depends on mysql

## run using docker compose
docker compose up -d

## run using docker file
docker pull mysql:latest
docker run --name mysqldb -e MYSQL_ROOT_PASSWORD=01emil01 -p 3306:3306 -d mysql

docker build --rm -t project-db-service . 
docker run -p 9022:9022 -d project-db-service