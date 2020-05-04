FROM postgres:9.6-alpine

COPY scripts/init.sql /docker-entrypoint-initdb.d/

COPY scripts/roleinit.sh /docker-entrypoint-initdb.d/

EXPOSE 5432