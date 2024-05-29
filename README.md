# Limit

A HTTP/HTTPS request rate limiter written in Go

## Getting Started

### Prerequisites

#### Environment Variables

You can set the following environment variables to configure the rate limiter:

- `FORWARD_URL` (**required**): The URL to forward requests to
- `RATE_LIMIT` (**required**): The number of requests allowed per minute
- `TIME_WINDOW` (**required**): The time window in seconds
- `PORT` (default: `7654`): The port the server will listen on
- `REDIS_ADDR` (default: `localhost:6379`): The address where the Redis server is running
- `REDIS_PASSWORD` (default: `""`): The password of the Redis server
- STILL TODO: `REDIS_DB` (default: `0`): The database of the Redis server
- `USE_TLS` (default: `false`): Whether to use TLS

You can store these variables in a `.env` file in the root of the project. If there is no `.env` file, the server will 
check to see if the variables are otherwise set. Variables stored in .env will take precedence over those set in the
environment.

**If you are using Docker**, you will need to set these variables in the `docker-compose.yml` file. 

#### Redis

The rate limiter uses Redis to store the rate limit data. 

The `docker-compose.yml` file in the root of the project contains a Redis service. If you are not using `docker compose`
you will need to set up a Redis server yourself.


