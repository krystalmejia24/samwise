# samwise

## Dependencies
 - [Docker](https://www.docker.com/products/docker-desktop)
 - [Redis](https://redis.io/download)

## Environment

| ENV VAR                 | Default VALUE       | Description              |
|-------------------------|:-------------------:|--------------------------|
| `SAMWISE_ENV`           | `local`             | environment              |
| `SAMWISE_HTTP_PORT`     | `80808`             | http server listen port  |
| `SAMWISE_TIMEOUT`       | `5s`                | http server timeout      |
| `SAMWISE_DB_CONNECTION` | `redis-server:6379` | Redis connection string. Uses Docker default image by default |

## Quick Start
1. Clone the repo
```
git clone git@github.com:krystalmejia24/samwise.git
```

2. Export all the proper [environment](#Environment) variables as described

3. Build and Run the service

To run w/ Docker. You can update the ports [here](https://github.com/krystalmejia24/samwise/blob/master/docker-compose.yml#L6):
```
make build_docker
make run_docker
```

To run the service locally, you will first need to set up a local instance of redis and export the proper enviornment variable for the DB connection. Once redis is setup, you can simply build and run the service with:
```
make build
make run
```

## Tests

tests only:
```
make test
```
coverage only:
```
make cover
```
tests & coverage:
```
make test_cover
```
benchmarks:
```
make bench
```