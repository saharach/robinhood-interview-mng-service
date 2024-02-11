#!/bin/bash

# Define SQL content
SQL_CONTENT=$(cat <<EOF
CREATE DATABASE $POSTGRES_DBNAME;
GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DBNAME TO $POSTGRES_USERNAME;
EOF
)

# Write SQL content to file
echo "$SQL_CONTENT" > /docker-entrypoint-initdb.d/init-script.sql

# Start PostgreSQL
exec docker-entrypoint.sh postgres