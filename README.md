# Limitr

A HTTP/HTTPS request rate limiter written in Go.

## About Limitr

Limitr works by forwarding requests to a specified URL and moderating the rate at which requests can be made in
accordance with a defined number of requests per defined time window. If the rate limit is exceeded, the server will
respond with a `429 Too Many Requests` status code.

### Why did I write Limitr?

I had a need for rate limiting a self-hosted API that I was developing, and couldn't find any standalone rate limiters
that weren't part of some API gateway or reverse proxy. I wanted something that I could run on my server and configure
to my liking.

### Technologies used

- [Go](https://golang.org/): The server is written in Go
- [Go Fiber](https://gofiber.io/): The server uses Go Fiber to handle HTTP requests
- [Redis](https://redis.io/): The rate limiter uses Redis to store rate limit data
- [Docker](https://www.docker.com/): The server can be run in a Docker container

#### Why Go and Why Redis?

I chose Go primarily for its speed. Since the server is a rate limiter, it needs to be able to handle requests quickly.
I chose Redis for the same reason. Redis is an in-memory database that is very fast and can handle a large number of
requests.

## Getting Started

### Prerequisites

#### Environment Variables

You can set the following environment variables to configure the rate limiter:

- `FORWARD_URL` (**required**): The URL to forward requests to
- `RATE_LIMIT` (**required**): The number of requests allowed per time window
- `TIME_WINDOW` (**required**): The time window in seconds
- `PORT` (**optional**, default: `7654`): The port the server will listen on
- `USE_TLS` (**optional**, default: `false`): Whether to use TLS (certificates are required)
- `IP_HEADER_KEY` (**optional**, default: blank): The header key that contains the client's IP address
- `REDIS_ADDR` (**optional**, default: `localhost:6379`): The address where the Redis server is running
- `REDIS_PASSWORD` (**optional**, default: `""`): The password of the Redis server. Please set a new password here and
  in `redisconfig/redis.conf` if you are using Redis in a production environment
- `REDIS_DB` (**optional**, default: `0`): The database of the Redis server
- `VERBOSE_MODE` (**optional**, default: `false`): Whether to print incoming requests to the console

You can store these variables in a `.env` file in the same directory as the executable. If there is no `.env` file, the
server will check to see if the variables are otherwise set. Variables stored in .env will take precedence over those
already set in the environment before running.

If you are using `docker compose`, you will need to set these variables in the `docker-compose.yml` file.

#### Redis

The rate limiter uses Redis to store the rate limit data.

The `docker-compose.yml` file in the root of the project contains a Redis service. If you are not using `docker compose`
you will need to set up a Redis server yourself. Make sure to set your environment variables appropriately (see the
[environment variables](#environment-variables) section).

#### Some Quick Info

The `master` branch contains the latest stable release. Active development is done in the `dev` branch.

### Running Limitr Normally

#### Using Docker Compose

**Note:** When running in Docker, Limitr cannot accurately get the client's IP address from the request. You will need
to run the container behind a reverse proxy (such as Nginx, HAProxy, Cloudflare Tunnels, etc.) that can forward the
client IP address in a header. You will then need to set the `IP_HEADER_KEY` environment variable to the header key that
contains the client's IP address. Alternatively, you can run Limitr without Docker, or you can
modify [docker-compose.yml](docker-compose.yml) to run both containers in host network mode (only available on Linux
hosts), both of which will allow you to get the client's IP address directly from the request. If you choose one of
these options, please make sure to secure your Redis server accordingly.

**Note:** The current Docker Compose configuration is set up to receive the client IP address from a header passed by
Cloudflare Tunnels. If you are not using Cloudflare Tunnels, you will need to modify the `docker-compose.yml` file to
suit your needs according to the note above.

First, clone this repository:

```shell
git clone https://github.com/BeehiveBroadband/limitr.git
cd limitr
```

To run the server using Docker Compose, you can use the following command from within the project directory root:

```shell
docker-compose up
```

This will start the Limitr server and a Redis server. The server will be available at `http://localhost:7654`.

**Note:** If you upgrade to a new release, you may need to delete the old limitr image from your Docker environment
before running `docker compose up`.

#### Using the executable

You can download the executable on the [releases page](https://github.com/BeehiveBroadband/limitr/releases). Download
the appropriate release for your system and run the executable.

```shell
./[executable name]
```

**Note:** You may need to make the downloaded file executable by running `chmod +x [file name]` on UNIX-like systems.

The server will be available at `http://localhost:7654` unless you set a different port.

##### Checking the version

You can check the version of the executable by running the following command:

```shell
./[executable name] -v
```

#### Building from source

You can build from source if a release isn't available for your system. Please note that you will need to have Go >=
1.22.3 installed on your system.

First, clone the repository:

```shell
git clone https://github.com/BeehiveBroadband/limitr.git
cd limitr
```

Then, build the executable:

```shell
go build ./cmd/limitr
```

Finally, run the executable:

```shell
./limitr
```

The server will be available at `http://localhost:7654` unless you set a different port.

#### Building from the `dev` branch (unstable)

If you want to build from the dev branch to get the latest features, you can do so. Please note that the dev branch may
contain unstable code. You will need to have Go >= 1.22.3 installed on your system to build from source.

First, clone the repository:

```shell
git clone https://github.com/BeehiveBroadband/limitr.git
cd limitr
```

Then, checkout the `dev` branch:

```shell
git checkout dev
```

Then, build the executable:

```shell
go build ./cmd/limitr
```

Finally, run the executable:

```shell
./limitr
```

The server will be available at `http://localhost:7654` unless you set a different port.

## Licensing

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

