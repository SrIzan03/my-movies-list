CREATE TABLE IF NOT EXISTS Account (
    id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    username varchar(100) NOT NULL UNIQUE,
    email varchar(254) NOT NULL UNIQUE,
    pwd_hash
)