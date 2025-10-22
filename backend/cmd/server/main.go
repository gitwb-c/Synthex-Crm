package main

import (
	"log"

	"github.com/gitwb-c/crm.saas/backend/internal/db"
	"github.com/gitwb-c/crm.saas/backend/internal/http"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	client, er := db.NewClient()
	if er != nil {
		log.Fatalf("failed to init Ent client: %v", er)
	}
	srv, err := db.Init(client)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer db.Close(client)

	r := http.GlobalRouter(srv)
	r.Run(":8080")
}
