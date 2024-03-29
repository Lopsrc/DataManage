-- CREATE TABLE IF NOT EXISTS auth
-- (
--     user_id     INTEGER REFERENCES user(id),
--     price       INTEGER NOT NULL
-- );

-- CREATE TABLE IF NOT EXISTS user
-- (
--     id              SERIAL PRIMARY KEY,
--     login           VARCHAR(255) NOT NULL UNIQUE,
--     pass_hash       bytea  NOT NULL,
--     del             BOOLEAN NOT NULL DEFAULT FALSE
-- );

CREATE TABLE IF NOT EXISTS prices
(
    user_id     SERIAL PRIMARY KEY,
    price       INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS work
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    date        DATE NOT NULL,
    price       INTEGER NOT NULL,
    time        INTEGER NOT NULL,
    penalty     INTEGER,
    user_id     INTEGER REFERENCES prices(user_id)
);
