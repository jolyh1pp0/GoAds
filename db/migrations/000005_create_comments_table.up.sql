CREATE TABLE comments (
    id SERIAL NOT NULL PRIMARY KEY,
    advertisement_id INT NOT NULL,
    content VARCHAR(1000) NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT Now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT Now(),

    CONSTRAINT fk_advertisement
      FOREIGN KEY(advertisement_id)
          REFERENCES advertisements(id)
          ON DELETE SET NULL,

    CONSTRAINT fk_user
      FOREIGN KEY(user_id)
          REFERENCES users(id)
);