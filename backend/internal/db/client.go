package db

import (
	"context"
	"log"

	"crm.saas/backend/internal/ent"
	"crm.saas/backend/internal/ent/migrate"
	"crm.saas/backend/internal/graphql/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/lib/pq"
)

var Client *ent.Client

func Init(connString string) (*handler.Server, error) {
	var err error
	Client, err = ent.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := Client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Client: Client}}))
	return srv, nil
}

func Close() {
	if Client != nil {
		Client.Close()
	}
}
