package handlers

import (
	"idontlikequizlet/internal/db"
	"idontlikequizlet/internal/tmpl"
	"net/http"
)

func HandleDecks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Pool.Query(r.Context(), "select id, name from decks order by id")
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

	data := map[string]any{
		"Message": "Decks from DB:",
		"Decks":   decks,
	}

	if err := tmpl.T.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 500)
	}
}