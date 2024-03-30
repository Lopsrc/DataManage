#!/bin/bash

sel=$1
# create db.

if [[ "$sel" == "local" ]]; then
    PGPASSWORD=postgres psql -h localhost -p 5432 -U postgres -c "IF EXISTS (SELECT FROM pg_database WHERE datname = 'datamanage') THEN
            RAISE NOTICE 'Database already exists';  -- optional
        ELSE
            PERFORM dblink_exec('dbname=' || current_database()  -- current db
                            , 'CREATE DATABASE datamanage');
        END IF;
    END"
elif [[ "$sel" == "docker" ]]; then
    psql -h db -p 5432 -U postgres
else
    echo "Fuck you."
fi