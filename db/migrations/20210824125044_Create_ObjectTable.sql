-- migrate:up
CREATE TABLE "object" (
    id SERIAL PRIMARY KEY,
    object_id INT,
    online BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_object UNIQUE (object_id)
);

-- migrate:down
DROP TABLE IF EXISTS "object";
