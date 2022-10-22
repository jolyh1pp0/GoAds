CREATE TABLE gallery (
    id SERIAL NOT NULL PRIMARY KEY,
    advertisement_id INT NOT NULL,
    file_path varchar(255) NOT NULL,
    file_name varchar(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT Now(),

    CONSTRAINT fk_advertisement
        FOREIGN KEY(advertisement_id)
            REFERENCES advertisements(id)
            ON DELETE SET NULL
);
