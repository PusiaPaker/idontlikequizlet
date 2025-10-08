idontlikequizlet
----------------

A small Go + PostgreSQL flashcard web app that mimics Quizlet-style decks,
but keeps everything server-rendered with templates and simple JavaScript.

STACK
-----
- Go 1.25 (using chi, pgx, godotenv)
- PostgreSQL
- HTML templates with Bootstrap 5
- Vanilla JS for client interactions

PROJECT STRUCTURE
-----------------
    .
    ├── cmd/server/main.go          - main entrypoint and router setup
    ├── internal/
    │   ├── db/                     - database connection and pool management
    │   ├── handlers/               - HTTP route handlers
    │   └── tmpl/                   - template loading and function map
    ├── migrations/                 - database schema and seed SQL files
    └── web/
        ├── static/                 - css and js assets
        └── templates/              - base templates and partials

QUICKSTART
-----------
1) Requirements:
   - Go 1.25 or higher
   - PostgreSQL running locally or remotely
   - DATABASE_URL environment variable set, for example:

       export DATABASE_URL="postgres://user:pass@localhost:5432/quizlet?sslmode=disable"

2) Install dependencies:
       go mod download

3) Initialize the database:
   Run the SQL files in the migrations folder in numerical order.
   Example using psql:

       psql "$DATABASE_URL" -f migrations/0001_create_users.sql
       psql "$DATABASE_URL" -f migrations/0002_create_user_stats.sql
       psql "$DATABASE_URL" -f migrations/0003_create_decks.sql
       psql "$DATABASE_URL" -f migrations/0004_create_cards.sql
       psql "$DATABASE_URL" -f migrations/0005_seed_users.sql
       psql "$DATABASE_URL" -f migrations/0006_seed_decks_and_cards.sql

   The seed files can be safely re-run since they use ON CONFLICT DO NOTHING.

4) Run the server:
       go run ./cmd/server

   The app will start on port 3000 by default.
   Visit http://localhost:3000 in your browser.

ROUTES
------
    GET    /                             homepage listing all decks
    GET    /deck/{deckID}                view a single deck and study mode
    GET    /edit/{deckID}                edit deck title and cards
    DELETE /edit/{deckID}/delete/{cardID} delete a card
    PATCH  /edit/{deckID}/update/{cardID} update card term or definition
    PATCH  /edit/{deckID}/update/title   rename the deck
    GET    /add                          create a new deck (for demo user)
    POST   /add/{deckID}                 add a blank card to an existing deck
    GET    /ping                         check database connectivity

NOTES
-----
- Seed users: pusiapaker, morgan, and allie
- The deck editor autosaves with debounce logic to reduce requests
- Blank terms and definitions are allowed for initial creation
- Static files are served under /static/*
- If you move templates or static folders, update tmpl.MustInit() and main.go

DEVELOPMENT TIPS
----------------
- Add chi middlewares such as RequestID, Recoverer, Logger, and Compress
- Consider adding a repository layer for database access (internal/store)
- Add graceful shutdown to the HTTP server
- Move inline CSS in templates to dedicated CSS files for caching
- Add created_at and updated_at timestamps and indexes to the database
- For deployment as a single binary, use Go’s //go:embed to bundle static assets

LICENSE
-------
MIT
