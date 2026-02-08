-- USERS TABLE
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,                    
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    career TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX IF NOT EXISTS users_email_unique_idx
ON users (LOWER(email));

CREATE INDEX IF NOT EXISTS users_career_idx
ON users (career);

CREATE TABLE IF NOT EXISTS user_interests (
    id BIGSERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    interest TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT fk_user_interests_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS user_interests_user_id_idx
ON user_interests (user_id);

CREATE UNIQUE INDEX IF NOT EXISTS user_interests_unique_idx
ON user_interests (user_id, interest);
