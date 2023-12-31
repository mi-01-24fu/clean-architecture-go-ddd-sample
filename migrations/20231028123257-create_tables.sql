
-- +migrate Up
CREATE TABLE users (
   id SERIAL PRIMARY KEY,
   username VARCHAR(255) NOT NULL,
   mailaddress VARCHAR(255) NOT NULL UNIQUE,
   password_hash VARCHAR(255) NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
   id SERIAL PRIMARY KEY,
   users_id INTEGER REFERENCES Users(id) ON DELETE CASCADE,
   title VARCHAR(255) NOT NULL,
   content TEXT,
   date DATE NOT NULL,
   status VARCHAR(255) CHECK (status IN ('未完了', '完了', '保留中')),
   priority VARCHAR(255) CHECK (priority IN ('高', '中', '低')),
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;
