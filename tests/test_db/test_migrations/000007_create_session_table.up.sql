CREATE TABLE sessions (
      id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),

      user_id UUID NOT NULL,
      access_token_uuid UUID NOT NULL,
      refresh_token_uuid UUID NOT NULL,

      refresh_token_expires_at TIMESTAMPTZ NOT NULL,
      expires_at TIMESTAMPTZ NOT NULL,

      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

      CONSTRAINT fk_user
          FOREIGN KEY(user_id)
              REFERENCES users(id)
);
