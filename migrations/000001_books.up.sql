CREATE TABLE IF NOT EXISTS books (
    id SERIAL NOT NULL UNIQUE PRIMARY KEY,
    name varchar(255) NOT NULL,
    author varchar(255) NOT NULL
);