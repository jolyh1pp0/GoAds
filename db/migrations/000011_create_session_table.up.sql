CREATE TABLE sessions (
    uuid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    access_token VARCHAR NOT NULL,
    access_token_uuid UUID NOT NULL,
    refresh_token VARCHAR NOT NULL,
    refresh_token_uuid UUID NOT NULL
);