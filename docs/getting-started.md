
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
There is a [default production configuration](../docker-compose.yml) which can be directly used for building and running the tool. This configuration can be modified based on your requirement as follows:

**`api-service` service:**

*Environment Variables:*

The following environment variables are used to configure the application:

- `RATE_LIMIT`: This variable specifies the maximum number of requests allowed per IP address within a certain time period.
  - Default value: `50`
- `RATE_LIMIT_TIME`: This variable sets the time period (in minutes) for the rate limit.
  - Default value: `1`
- `POSTGRES_HOST`: This variable specifies the hostname of the PostgreSQL database server.
  - Default value: `database`
- `POSTGRES_PORT`: This variable specifies the port number on which the PostgreSQL database server is running.
  - Default value: `5432`
- `POSTGRES_USERNAME`: This variable specifies the username used to authenticate with the PostgreSQL database server.
  - Default value: `postgres`
- `POSTGRES_PASSWORD`: This variable specifies the password used to authenticate with the PostgreSQL database server.
  - Default value: `password`
- `POSTGRES_DBNAME`: This variable specifies the name of the PostgreSQL database used by the application.
  - Default value: `interview`

**`mysql` service:**

*Environment Variables:*

- `POSTGRES_USERNAME`: Specifies the username used to authenticate with the PostgreSQL database server.
  - Default value: `postgres`
- `POSTGRES_PASSWORD`: Specifies the password used to authenticate with the PostgreSQL database server.
  - Default value: `password`
- `POSTGRES_DBNAME`: Specifies the name of the PostgreSQL database used by the application.
  - Default value: `interview`
*Volumes:*

PostgreSQL data is stored in `/var/lib/postgresql/data` folder inside docker container and mounted to `./postgresql/data` host server volume.


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
## Troubleshooting Docker Start Error

If you encounter errors when starting Docker containers due to permission issues, you can resolve them by adjusting the file permissions using the `chmod` command.

### Fixing Permissions

1. **Identify the problematic file**: 
   - Determine which file or directory is causing the permission error.

2. **Adjust permissions using `chmod`**:
   - Open a terminal or command prompt.

   - Run the following command to give read, write, and execute permissions to the file:
     ```sh
     chmod +x /path/to/your/file
     ```

   Replace `/path/to/your/file` with the path to the file or directory causing the permission error.

3. **Restart Docker containers**:
   - After adjusting the permissions, restart your Docker containers to apply the changes:
     ```sh
     docker-compose restart
     ```

### Example
If you encounter a permission error with the `init-database.sh` script, you can fix it by running:
```sh
chmod +x ./postgresql/init-database.sh
```


## Initial User Credentials

To access the API for the first time, you'll need to use the initial user credentials provided below:

**Admin role:**
- **Username**: `admin`
- **Password**: `password`

**User role:**
- **Username**: `user1`
- **Password**: `password`

Please note that these credentials are for initial access only. 


## Postman API Documentation

We've provided API documentation using Postman to help you interact with our API endpoints effortlessly. Follow the instructions below to import the collection into Postman:

1. **Download the Collection**: 
   - Download the Postman Collection JSON file from [this link](./api-document/Robinhood.postman_collection.json)

2. **Import into Postman**:
   - Open Postman and click on the "Import" button located in the top left corner.
   - Select the downloaded JSON file and click "Open" to import the collection into Postman.

3. **Explore the API**:
   - Once imported, you'll see a collection named "API Documentation" containing all the requests and documentation.
   - Explore the requests to understand the available endpoints, their functionalities, and the required parameters.
   - Each request may include detailed descriptions, examples, and documentation to guide you through the API usage.

By following these steps, you'll have access to our API documentation in Postman, making it easy to interact with our API and integrate it into your applications. If you have any questions or need assistance, feel free to reach out to us.




