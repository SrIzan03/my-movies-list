package main

import (
	"back/cmd/api-server/routes"
	"back/internal/auth"
	"back/internal/config"
	"back/internal/db"
	"context"
	"database/sql"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	conn, err := db.ConnectToDB(cfg.DB, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	mux := http.NewServeMux()
	authService := auth.NewService(cfg, ctx)

	routes.RegisterRoutes(mux, authService)
}
