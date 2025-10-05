CREATE TABLE IF NOT EXISTS cards (
  id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  deck_id     uuid NOT NULL REFERENCES decks(id) ON DELETE CASCADE,
  term        text NOT NULL,
  definition  text NOT NULL,
  UNIQUE(deck_id, term)
);