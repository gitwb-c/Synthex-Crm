package main

import (
	"log"
	"os"

	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/http"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}
	srv, err := db.Init(os.Getenv("DB_POSTGRES_URL"))
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer db.Close()

	r := http.GlobalRouter(srv)
	r.Run(":8080")

}
