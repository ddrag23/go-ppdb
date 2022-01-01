CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY,
    photo VARCHAR(255) null,
    name VARCHAR (255) null,
    email VARCHAR (255) null,
    username VARCHAR (255) unique null,
    password VARCHAR (255) null,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP default CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP null
);

CREATE TABLE IF NOT EXISTS roles(
    id uuid PRIMARY KEY,
    name VARCHAR (50) null,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP default CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP null
);