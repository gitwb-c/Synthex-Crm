package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"crm.saas/backend/internal/db"
	"crm.saas/backend/internal/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}
	if err := db.Init(os.Getenv("DB_POSTGRES_URL")); err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer db.Close()

	http.GlobalRouter()
}
