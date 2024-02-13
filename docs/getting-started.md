
# Getting started

## Requirements


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

### Production

#### Configuration
There is a [default production configuration](../docker-compose.prod.yml) which can be directly used for building and running the tool. This configuration can be modified based on your requirement as follows:

**`backend` service:**

*Environment Variables:*

1. `ADMIN_USERNAME`: Username for admin user (defaults to `admin`)
2. `ADMIN_PASSWORD`: Password for admin user (defaults to `password`)
3. `DATABASE_URL`: SQLAlchemy Database URL (currently only MySQL database is supported)
4. `JWT_SECRET_KEY`: JSON Web Token Secret key
5. `JWT_REDIS_STORE_URL`: JSON Web Token Redis Store URL

*Volumes:*

Audio datapoints uploaded are stored in `/root/uploads` folder inside docker container and mounted to `backend_data` volume. You can change this and mount host server volume instead.

**`mysql` service:**

*Environment Variables:*

1. `MYSQL_DATABASE`: MySQL Database name. Defaults to `audino`. If changed, you need to change database name in `../mysql/create_database.sql`.
2. `MYSQL_ROOT_PASSWORD`: Password for `root` user. Defaults to `root`.
3. `MYSQL_USER`: Application user to be created for `MYSQL_DATABASE`. *Note: `DATABASE_URL` in `backend` service should reflect this change*
4. `MYSQL_PASSWORD`: Application user's password. *Note: `DATABASE_URL` in `backend` service should reflect this change*

*Volumes:*

MySQL data is stored in `/var/lib/mysql` folder inside docker container and mounted to `mysql_prod_data` volume. You can change this and mount host server volume instead.

**`redis` service:**

*Environment Variables:*

1. `REDIS_PASSWORD`: Password for redis store. Defaults to `audino`. *Note: `JWT_REDIS_STORE_URL` in `backend` service should reflect this change*

*Volumes:*

Redis data is stored in `/data` folder inside docker container and mounted to `redis_data` volume. You can change this and mount host server volume instead.


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

## Tutorials

We provide a set of [tutorials](./docs/tutorial.md) to guide users to achieve certain tasks.
