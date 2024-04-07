CREATE TABLE IF NOT EXISTS users (
    uuid UUID PRIMARY KEY,
    login VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    encrypted_password VARCHAR NOT NULL,
    is_admin BOOLEAN DEFAULT FALSE
);