#!/usr/local/bin/bash

# - find all directories named schema
# - execute the model generation code on them

services=(
    "chat"
)

function gen() {
    local _info="${1^^}_POSTGRES_INFO"
    local _db=$(echo ${!_info} | tr ',' ' ' | awk '{print $1}') 
    local _un=$(echo ${!_info} | tr ',' ' ' | awk '{print $2}') 

    mkdir -p "${1}/models"
    echo xo "pgsql://${_un}:${POSTGRES_PASSWORD}@0.0.0.0:5432/${_db}?sslmode=disable" -o "${1}/models"
}

# export all environment variables
export $(egrep -v '^#' .env | xargs)

# check and see if there was a specific service requested
if [[ $# -eq 1 ]]; then
    if [[ ! " ${services[@]} " =~ " ${1} " ]]; then
        echo "Service ${1} does not exist"
        exit 1
    fi
    echo "Generating code for service: ${1}"
    gen $1
    exit 0
fi

# if nothing was specified, then we can generate for all
for svc in "${services[@]}"; do
    echo "Generating code for service: ${svc}"
    gen $svc || exit 1
done