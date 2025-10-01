-- add a deck
insert into decks (name) values ('Russian 1 – Chapter 1') returning id;
-- assume it gave id=1

-- add cards
insert into cards (deck_id, front, back)
values
  (1, 'Здравствуйте', 'Hello'),
  (1, 'Как дела?', 'How are you?')
returning id;

insert into reviews (card_id, due_at) values
  (1, now()),
  (2, now());
