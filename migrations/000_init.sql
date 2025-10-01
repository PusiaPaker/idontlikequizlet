-- Decks = collections of flashcards
create table decks (
  id bigserial primary key,
  name text not null,
  created_at timestamptz default now()
);

-- Cards = front/back text
create table cards (
  id bigserial primary key,
  deck_id bigint references decks(id) on delete cascade,
  front text not null,
  back text not null,
  created_at timestamptz default now()
);

-- Reviews = spaced repetition state for each card
create table reviews (
  id bigserial primary key,
  card_id bigint unique references cards(id) on delete cascade,
  due_at timestamptz not null,
  interval_days int default 0,
  ease double precision default 2.5,
  streak int default 0,
  last_grade int
);
