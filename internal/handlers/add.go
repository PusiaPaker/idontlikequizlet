package handlers

import (
	"encoding/json"
	"idontlikequizlet/internal/db"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HandleAddCard(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	deckID := chi.URLParam(r, "deckID")

	type newCard struct {
		ID			string `json:"id"`
		Term		string `json:"term"`
		Definition 	string `json:"definition"`
	}

	var c newCard
	err := db.Pool.QueryRow(ctx, `
		INSERT INTO cards (deck_id, term, definition)
		VALUES ($1, '', '')
		RETURNING id, term, definition
	`, deckID).Scan(&c.ID, &c.Term, &c.Definition)
	if err != nil {
		http.Error(w, "could not create card", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(c)
}

func HandleCreateDeck(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	// TODO: Replace with real user ID from session once you add auth
	var ownerID string
	err := db.Pool.QueryRow(ctx, `SELECT id FROM users WHERE username = $1`, "pusiapaker").Scan(&ownerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var deckID string
	err = db.Pool.QueryRow(ctx, `
		INSERT INTO decks (owner_id, name)
		VALUES ($1, 'Untitled Deck')
		RETURNING id
	`, ownerID).Scan(&deckID)
	if err != nil {
		log.Printf("failed to create deck: %v", err)
		http.Error(w, "could not create deck", http.StatusInternalServerError)
		return
	}

	// Create one blank card so the user immediately sees a card editor box
	_, err = db.Pool.Exec(ctx, `
		INSERT INTO cards (deck_id, term, definition)
		VALUES ($1, '', '')
	`, deckID)
	if err != nil {
		log.Printf("failed to create first card: %v", err)
	}

	http.Redirect(w, r, "/edit/"+deckID, http.StatusSeeOther)
}