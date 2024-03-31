#!/bin/bash

sel=$1
hostremote=db
localhost=localhost
# create db.

if [[ "$sel" == "local" ]]; then
    PGPASSWORD=postgres psql -h "$localhost" -p 5432 -U postgres -c "SELECT 'CREATE DATABASE datamanage'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'datamanage')"
    PGPASSWORD=postgres psql -h "$localhost" -p 5432 -U postgres -d datamanage -a -f "server/migrations/1_init.up.sql"
elif [[ "$sel" == "docker" ]]; then
    PGPASSWORD=postgres psql -h "$hostremote" -p 5432 -U postgres -c "SELECT 'CREATE DATABASE datamanage'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'datamanage')"
    PGPASSWORD=postgres psql -h "$hostremote" -p 5432 -U postgres -d datamanage -a -f "server/migrations/1_init.up.sql"
else
    echo "Invalid parameters.
Example: ./migrate.sh local"
fi