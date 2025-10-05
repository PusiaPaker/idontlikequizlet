package handlers

import (
	"bytes"
	"idontlikequizlet/internal/db"
	"idontlikequizlet/internal/tmpl"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type card struct {
	ID			string		`json:"id"`
	Term		string		`json:"term"`
	Definition	string		`json:"definition"`
}

func HandleDeck(w http.ResponseWriter, r *http.Request) {
	deckID := chi.URLParam(r, "deckID")
	ctx := r.Context()

	// Asks Databse for the Deck Name
	var deckName string
	err := db.Pool.QueryRow(ctx, `SELECT name FROM decks WHERE id = $1`, deckID).Scan(&deckName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Asks DB for cards inside the deck
	rows, err := db.Pool.Query(ctx, `SELECT id, term, definition FROM cards WHERE deck_id = $1 ORDER BY id`, deckID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cards []card
	for rows.Next() {
		var c card
		if err := rows.Scan(&c.ID, &c.Term, &c.Definition); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cards = append(cards, c)
	}

	data := struct {
		Active string
		DeckName string
		Cards []card 
	} {
		Active: "home",
		DeckName: deckName,
		Cards: cards,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var buf bytes.Buffer
	if err := tmpl.Deck.ExecuteTemplate(&buf, "deck", data); err != nil {
		log.Printf("template exec error: %v", err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
	if _, err := buf.WriteTo(w); err != nil {
		log.Printf("write error: %v", err)
	}
}
