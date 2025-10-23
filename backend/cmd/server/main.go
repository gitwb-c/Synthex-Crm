package main

import (
	"log"

	"github.com/gitwb-c/crm.saas/backend/internal/db"
	"github.com/gitwb-c/crm.saas/backend/internal/http"
	"github.com/gitwb-c/crm.saas/backend/pkg/env"
)

func main() {
	err := env.InitEnv()
	if err != nil {
		return
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

	r := http.GlobalRouter(client, srv)
	r.Run(":8080")
}
