package main

import (
	"log"

	"github.com/gitwb-c/crm.saas/backend/internal/db"
	"github.com/gitwb-c/crm.saas/backend/internal/http"
)

func main() {
	client, _ := db.NewClient()
	srv, err := db.Init(client)
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	defer db.Close(client)

	r := http.GlobalRouter(srv)
	r.Run(":8080")

}
