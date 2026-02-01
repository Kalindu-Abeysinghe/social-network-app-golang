package main

import (
	"log"

	"github.com/Kalindu-Abeysinghe/social-app.git/internal/db"
	"github.com/Kalindu-Abeysinghe/social-app.git/internal/env"
	"github.com/Kalindu-Abeysinghe/social-app.git/internal/store"
)

func main() {
	cfg := config{
		address: env.GetString("SERVER_ADDR", "localhost:8085"),
		db: dbConfig{
			addr: env.GetString("DB_ADRR", "postgres://postgres:Gyarados@localhost:5433/socialnetwork?sslmode=disable"),
			// addr:               env.GetString("DB_ADRR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConnections: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConnections: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:        env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env:     env.GetString("ENV", "local"),
		version: env.GetString("VERSION", "0.0.0"),
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConnections,
		cfg.db.maxOpenConnections,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Database connection established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	appErr := app.run(mux)

	log.Fatal(appErr)
}
