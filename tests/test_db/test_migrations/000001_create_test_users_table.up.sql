CREATE TYPE verified_types AS ENUM('email', 'phone', 'google');
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    time_zone VARCHAR(255) NOT NULL,
    phone VARCHAR(255) UNIQUE,
    disabled BOOLEAN NOT NULL DEFAULT false,
    verified_type verified_types,
    google_secret VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)