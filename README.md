# Rate Limiter (Go)
[![Go](https://github.com/teootoledo/go-rate-limiter/actions/workflows/go.yml/badge.svg)](https://github.com/teootoledo/go-rate-limiter/actions/workflows/go.yml)

## Description
This is a Rate-Limited Notification Service that allows users to send notifications to their clients.
This service limits the number of notifications that can be sent to a client in a given time window for a specific mail type.

- Status: not more than 2 per minute for each recipient
- News: not more than 1 per day for each recipient
- Marketing: not more than 3 per hour for each recipient

## Tests

### Coverage
To run the tests and check the coverage, you can run the following command:
```bash
./bin/coverage.sh
```
Make sure to give execution permissions to the file:
```bash
chmod +x bin/coverage.sh
```

### Integration tests
This project includes Postman integration tests. To run them, you need to have Postman installed and import the collection located in the `local-testing` folder.

Then run the collection as `Functional`, `Run manually` and `Iterations` set to `1`, and `Delay` set to `0`.

## Run the API
### Local
#### Without Docker
To run the API, you can use the following command:
```bash
go run main.go
```

#### With Docker
To run the API with Docker, you can use the following command:
```bash
docker-compose up -d
```

## Documentation (Open API Specification)
The documentation of the API can be found [here](http://localhost:8080/v1/docs/index.html) if you run it locally or with docker.

## Contact
If you have any questions, feel free to contact me vía [email](mailto:teootoledo@gmail.com) or [LinkedIn](https://www.linkedin.com/in/teootoledo/).
