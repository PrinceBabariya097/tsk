-- +goose Up
ALTER TABLE todos
ADD COLUMN priority TEXT NOT NULL CHECK(priority IN ('low', 'medium', 'high')) DEFAULT 'medium';

ALTER TABLE todos
ADD COLUMN tag TEXT; 

-- +goose Down
ALTER TABLE todos
DROP COLUMN priority;

ALTER TABLE todos
DROP COLUMN tag;
