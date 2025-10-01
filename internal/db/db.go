package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MustConnect() *pgxpool.Pool {
	url := "postgres://admin:postgresql@192.168.0.92:5432/idontlikequizlet"
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}
	return pool
}
