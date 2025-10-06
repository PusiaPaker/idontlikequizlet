package main

import (
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

	tmpl.MustInit()

	db.Pool = db.MustConnect()
	defer db.Pool.Close()

	r := chi.NewRouter()
	r.Get("/", handlers.HomeHandle)
	r.Get("/ping", handlers.HandlePing)
	r.Get("/deck/{deckID}", handlers.HandleDeck)
	
	r.Get("/edit/{deckID}", handlers.HandleEdit)
	r.Delete("/edit/{deckID}/delete/{cardID}", handlers.HandleDeleteCard)
	r.Patch("/edit/{deckID}/update/{cardID}", handlers.HandleUpdateCard)
	r.Patch("/edit/{deckID}/image/{cardID}", handlers.HandleUpdateCardImage)
	r.Patch("/edit/{deckID}/update/title", handlers.HandleUpdateTitle)
	
	r.Get("/add", handlers.HandleCreateDeck)
	r.Post("/add/{deckID}", handlers.HandleAddCard)


	// static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("listening on :3000")
	http.ListenAndServe(":3000", r)
}
