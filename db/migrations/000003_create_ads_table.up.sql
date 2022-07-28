CREATE TABLE advertisements (
    id SERIAL NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    title VARCHAR(200) NOT NULL UNIQUE,
    description VARCHAR(1000),
    price BIGINT NOT NULL,
    photo_1 VARCHAR(255),
    photo_2 VARCHAR(255),
    photo_3 VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT Now(),

    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
);