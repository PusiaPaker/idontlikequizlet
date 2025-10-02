package handlers

import (
	"net/http"

	"idontlikequizlet/internal/tmpl"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Active": "home",
		"Decks":  []struct{ ID int; Name string }{}, // fill from DB
	}
	if err := tmpl.T.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 500)
	}
}