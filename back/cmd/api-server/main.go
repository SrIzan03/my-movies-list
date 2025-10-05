package main

import (
	"back/internal/config"
	"back/internal/db"
	"database/sql"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	conn, err := db.ConnectToDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
}
