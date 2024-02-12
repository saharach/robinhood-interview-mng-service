#!/bin/bash

SQL_CONTENT=$(cat <<EOF
CREATE DATABASE $POSTGRES_DBNAME;
GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DBNAME TO $POSTGRES_USERNAME;
EOF
)

echo "$SQL_CONTENT" > /docker-entrypoint-initdb.d/init-script.sql

exec docker-entrypoint.sh postgres