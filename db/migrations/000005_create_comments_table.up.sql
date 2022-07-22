CREATE TABLE comments (
    id SERIAL NOT NULL PRIMARY KEY,
    advertisement_id INT NOT NULL,
    content VARCHAR(1000) NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT Now()
);