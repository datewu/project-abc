CREATE TABLE records (
    id SERIAL PRIMARY KEY,
    name text UNIQUE NOT NULL,
    status INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX pgroonga_videos_name_index ON videos USING pgroonga (name);

