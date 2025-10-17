package db

import (
	"context"
	"log"

	"crm.saas/backend/internal/ent"
	_ "github.com/lib/pq"
)

var Client *ent.Client

func Init(connString string) error {
	var err error
	log.Print(connString)
	Client, err = ent.Open("postgres", connString)
	if err != nil {
		return err
	}

	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return nil
}

func Close() {
	if Client != nil {
		Client.Close()
	}
}
