package main

import (
	"log"

	"github.com/Kalindu-Abeysinghe/social-app.git/internal/db"
	"github.com/Kalindu-Abeysinghe/social-app.git/internal/env"
	"github.com/Kalindu-Abeysinghe/social-app.git/internal/store"
)

func main() {
	conn, err := db.New(
		env.GetString("DB_ADRR", "postgres://postgres:Gyarados@localhost:5433/socialnetwork?sslmode=disable"),
		env.GetInt("DB_MAX_OPEN_CONNS", 30),
		env.GetInt("DB_MAX_IDLE_CONNS", 30),
		env.GetString("DB_MAX_IDLE_TIME", "15m"),
	)
	if err != nil {
		log.Panic(err)
	}

	err = conn.Close()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Database connection established")

	dbStore := store.NewStorage(conn)

	db.Seed(dbStore)
}
