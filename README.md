# todo Assessment

This project contains rest-service that listens on localhost:9090 for POST requests on /callback and gets the object information for every incoming object_id and
filter the objects by their "online" status.

According to task steps, the below steps should be run.

## Requirements

- RabbitMQ
- dbmate

## Steps to start

1. Clone the repo
2. Install dbmate [doc](https://github.com/amacneil/dbmate)
3. Install RabbitMQ [doc](https://www.rabbitmq.com/download.html)
4. Run [Steps to run](#steps-to-run)

## Steps to run

1. Go to main directory:
   ```
   cd todo/
   ```
2. Download `go` dependencies:
   ```
   go mod download
   ```
3. Run migration file:
   ```
   dbmate up
   ```
4. Run `callback` service:
   ```
   go run cmd/main.go run-api
   ```
5. Run `consumer` service:
   ```
   go run cmd/main.go run-consumer
   ```
6. Run test service:
   ```
   go run tester_service.go
   ```
 ## To improve
![architecture](https://freepngimg.com/thumb/street_fighter/35134-8-street-fighter-ii-image.png)