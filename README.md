
# interview management service

## Features

Current features of the API include:

1. JWT based authentication
2. Rate limit
3. Manage interviews

## Getting Started

For detailed instructions on getting started with service, please refer to the [Getting Started](./docs/getting-started.md) guide.


## Usage


Please install the following dependencies to run `service` on your system:

1. [git](https://git-scm.com/) *[tested on v2.42.0]*
2. [docker](https://www.docker.com/) *[tested on v25.0.2]*
3. [docker-compose](https://docs.docker.com/compose/) *[tested on v2.24.3]*

### Clone the repository

```sh
$ git clone https://github.com/saharach/robinhood-interview-mng-service.git
$ cd robinhood-interview-mng-service
```

**Note for Windows users**: Please configure git to handle line endings correctly as services might throw an error and not come up. You can do this by cloning the project this way:

```sh
$ git clone https://github.com/saharach/robinhood-interview-mng-service.git --config core.autocrlf=input
```

## Deploy

You can either run the project on [default configuration](./docker-compose.yml) or modify them to your need.
**Note**: Before proceeding further, you might need to give docker `sudo` access or run the commands listed below as `sudo`.

**To build the services, run:**

```sh
$ docker-compose  build
```

**To bring up the services, run:**

```sh
$ docker-compose  up -d
```

Then, test service via Postman  [http://0.0.0.0:8080/](http://0.0.0.0:8080/)

**To bring down the services, run:**

```sh
$ docker-compose down -v
```
