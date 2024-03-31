#!/bin/bash

sel=$1

hostremote=db
localhost=localhost

if [[ "$sel" == "local" ]]; then
    PGPASSWORD=postgres psql -h "$localhost" -p 5432 -U postgres -c "SELECT 'CREATE DATABASE datamanagertest'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'datamanagertest')"
    PGPASSWORD=postgres psql -h "$localhost" -p 5432 -U postgres -d datamanagertest -a -f "server/tests/migrations/1_init.up.sql"
elif [[ "$sel" == "docker" ]]; then
    PGPASSWORD=postgres psql -h "$hostremote" -p 5432 -U postgres -c "SELECT 'CREATE DATABASE datamanagertest'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'datamanagertest')"
    PGPASSWORD=postgres psql -h "$hostremote" -p 5432 -U postgres -d datamanagertest -a -f "server/tests/migrations/1_init.up.sql"
else
    echo "Invalid parameters.
Example: ./migrate.sh local"
fi