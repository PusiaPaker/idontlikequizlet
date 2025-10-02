package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"

	"idontlikequizlet/internal/db"
	"idontlikequizlet/internal/handlers"
	"idontlikequizlet/internal/tmpl"
)


func main() {
	_ = godotenv.Load()

	tmpl.T = template.New("").Funcs(nil)
	template.Must(tmpl.T.ParseGlob("web/templates/*.html"))
	template.Must(tmpl.T.ParseGlob("web/templates/partials/*.html"))

	db.Pool = db.MustConnect()
	defer db.Pool.Close()

	r := chi.NewRouter()
	r.Get("/", handlers.HomeHandle)
	r.Get("/ping", handlers.HandlePing)
	r.Get("/decks", handlers.HandleDecks)

	// static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("listening on :3000")
	http.ListenAndServe(":3000", r)
}
