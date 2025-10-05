
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE IF NOT EXISTS decks (
  id        uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  owner_id  uuid NOT NULL REFERENCES users(id) ON DELETE SET NULL,
  name      VARCHAR(128) NOT NULL,

  UNIQUE(owner_id, name)
);