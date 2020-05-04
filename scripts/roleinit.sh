#!/bin/bash

set -u

function create_user_and_database() {
	local database=$(echo $1 | tr ',' ' ' | awk  '{print $1}')
	local owner=$(echo $1 | tr ',' ' ' | awk  '{print $2}')
	local pass=$(echo $1 | tr ',' ' ' | awk  '{print $3}')
	echo "  Creating user and database '$database'"
	psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
	    CREATE USER $owner WITH ENCRYPTED PASSWORD '$pass';
	    CREATE DATABASE $database;
	    GRANT ALL PRIVILEGES ON DATABASE $database TO $owner;
EOSQL
}


echo "Multiple database creation requested"
for var in $(env | grep -o '^.*_POSTGRES_INFO'); do
    create_user_and_database ${!var}
done
echo "Multiple databases created"
