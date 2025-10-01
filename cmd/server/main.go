package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"idontlikequizlet/internal/db"
)

func main() {
	// connect to DB
	pool := db.MustConnect()
	defer pool.Close()

	// load templates
	tpl := template.Must(template.ParseGlob("web/templates/*.html"))

	// setup router
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "home.html", map[string]any{
			"Message": "Hello world",
		})
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		var now time.Time
		if err := pool.QueryRow(r.Context(), "select now()").Scan(&now); err != nil {
			http.Error(w, "db error: "+err.Error(), 500); return
		}
		w.Write([]byte("pong from DB, time: " + now.UTC().Format(time.RFC3339)))
	})

	r.Get("/decks", func(w http.ResponseWriter, r *http.Request) {
		rows, err := pool.Query(r.Context(), "select id, name from decks order by id")
		if err != nil {
			http.Error(w, "db error: "+err.Error(), 500)
			return
		}
		defer rows.Close()

		type deck struct {
			ID   int
			Name string
		}
		var decks []deck
		for rows.Next() {
			var d deck
			rows.Scan(&d.ID, &d.Name)
			decks = append(decks, d)
		}

		tpl.ExecuteTemplate(w, "home.html", map[string]any{
			"Message": "Decks from DB:",
			"Decks":   decks,
		})
	})

	// static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("listening on :3000")
	http.ListenAndServe(":3000", r)
}
