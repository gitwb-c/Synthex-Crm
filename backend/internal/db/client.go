package db

import (
	"context"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/ent/migrate"
	"github.com/gitwb-c/crm.saas/backend/internal/graphql/graph"
	_ "github.com/lib/pq"
)

var Client *ent.Client

func NewClient() (*ent.Client, error) {
	Client, err := ent.Open("postgres", os.Getenv("DB_POSTGRES_URL"))
	if err != nil {
		return nil, err
	}
	// Client.Use(hooks.TenantFilterHook)

	return Client, nil

}

func Init(client *ent.Client) (*handler.Server, error) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Client: client}}))
	return srv, nil
}

func Close(client *ent.Client) {
	if client != nil {
		client.Close()
	}
}
