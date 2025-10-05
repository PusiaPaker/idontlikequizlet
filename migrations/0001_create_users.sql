CREATE EXTENSION IF NOT EXISTS pgcrypto;

create table if not exists users (
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username      VARCHAR(32) UNIQUE NOT NULL,
  password      TEXT NOT NULL
);