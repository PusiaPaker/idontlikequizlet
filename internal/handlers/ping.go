package handlers

import (
	"net/http"
	"time"

	"idontlikequizlet/internal/db"
)

func HandlePing (w http.ResponseWriter, r *http.Request) {
	var now time.Time
	if err := db.Pool.QueryRow(r.Context(), "select now()").Scan(&now); err != nil {
		http.Error(w, "db error: "+err.Error(), 500); return
	}
	w.Write([]byte("pong from DB, time: " + now.UTC().Format(time.RFC3339)))
}
