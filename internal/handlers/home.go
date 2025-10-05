package handlers

import (
	"bytes"
	"idontlikequizlet/internal/db"
	"idontlikequizlet/internal/tmpl"
	"log"
	"net/http"
)

type deck struct {
	ID			string		`json:"id"`
	Name		string		`json:"name"`
}

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rows, err := db.Pool.Query(ctx, `SELECT id, name FROM decks`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	
	var decks []deck
	for rows.Next() {
		var d deck
		if err := rows.Scan(&d.ID, &d.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		decks = append(decks, d)
	}

	data := struct {
		Active string
		Decks []deck 
	} {
		Active: "home",
		Decks: decks,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var buf bytes.Buffer
	if err := tmpl.Home.ExecuteTemplate(&buf, "home", data); err != nil {
		log.Printf("template exec error: %v", err)
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
	if _, err := buf.WriteTo(w); err != nil {
		log.Printf("write error: %v", err)
	}
}
