package handlers

import (
	"bytes"
	"encoding/json"
	"idontlikequizlet/internal/db"
	"idontlikequizlet/internal/tmpl"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func HandleDeleteCard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	deckID := chi.URLParam(r, "deckID")
	cardID := chi.URLParam(r, "cardID")

	delQuery := `DELETE FROM cards WHERE (deck_id = $1 AND id = $2)`
	res, err := db.Pool.Exec(ctx, delQuery, deckID, cardID)
	if err != nil {
		log.Fatalf("Error executing DELETE query: %v\n", err)
		http.Error(w, "delete failed", http.StatusInternalServerError)
		return
	}

	log.Printf("Deleted %v row(s)\n", res.RowsAffected())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

func HandleUpdateCard(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	deckID := chi.URLParam(r, "deckID")
	cardID := chi.URLParam(r, "cardID")

	type CardPayload struct {
		Term       	*string `json:"term,omitempty"`
		Definition 	*string `json:"definition,omitempty"`
	}

	var p CardPayload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	if p.Term != nil {
	trimmed := strings.TrimSpace(*p.Term)
	p.Term = &trimmed
	}
	if p.Definition != nil {
		trimmed := strings.TrimSpace(*p.Definition)
		p.Definition = &trimmed
	}

	// Update only provided fields
	_, err := db.Pool.Exec(ctx, `
		UPDATE cards
		SET term = COALESCE($1, term),
		    definition = COALESCE($2, definition)
		WHERE deck_id = $3 AND id = $4
	`, p.Term, p.Definition, deckID, cardID)
	if err != nil {
		log.Println("update error:", err)
		http.Error(w, "update failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

func HandleUpdateTitle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	deckID := chi.URLParam(r, "deckID")

	type DeckTitlePayload struct {
		Name *string `json:"name"`
	}

	var p DeckTitlePayload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}
	if p.Name == nil {
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}
	title := strings.TrimSpace(*p.Name)

	_, err := db.Pool.Exec(ctx, `UPDATE decks SET name = $1 WHERE id = $2`, title, deckID)
	if err != nil {
		http.Error(w, "update failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}

func HandleUpdateCardImage(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

}

func HandleEdit(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	deckID := chi.URLParam(r, "deckID")
	
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
		Active 	string
		Name 	string
		DeckID	string
		Cards	[]card
	} {
		Active: "add",
		Name: 	deckName,
		DeckID: deckID,
		Cards: 	cards,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var buf bytes.Buffer
	if err := tmpl.Edit.ExecuteTemplate(&buf, "edit", data); err != nil {
		log.Printf("template exec error: %v", err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
	if _, err := buf.WriteTo(w); err != nil {
		log.Printf("write error: %v", err)
	}
}