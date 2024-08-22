package main

import (
	"log"
	"net/http"
	"webhooks/internal/config"
	"webhooks/internal/handler"
	"webhooks/internal/storage"
)

func main() {
	cfg := config.New()

	db, err := storage.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}
	defer db.Close()

	webhookHandler := handler.NewWebhookHandler(db)

	http.HandleFunc("/webhook", webhookHandler.Handle)

	log.Printf("Server started on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
