select * from decks;
select * from cards;
select * from reviews;

SELECT *
FROM cards
WHERE deck_id IN (
  SELECT id
  FROM decks
  WHERE owner_id = (SELECT id FROM users WHERE username = 'pusiapaker')
);
