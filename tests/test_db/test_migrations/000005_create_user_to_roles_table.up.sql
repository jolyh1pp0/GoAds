CREATE TABLE user_to_roles (
   id SERIAL NOT NULL PRIMARY KEY,
   user_id UUID NOT NULL,
   role_id INT NOT NULL,

   CONSTRAINT fk_user
       FOREIGN KEY(user_id)
           REFERENCES users(id),

   CONSTRAINT fk_role
       FOREIGN KEY(role_id)
           REFERENCES roles(id)
);