package db

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gitwb-c/crm.saas/backend/internal/ent"
	"github.com/gitwb-c/crm.saas/backend/internal/graphql/graph"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewClient() (*ent.Client, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}
	return ent.Open("postgres", os.Getenv("DB_POSTGRES_URL"))

}

func Init(client *ent.Client) (*handler.Server, error) {
	client, err := NewClient()
	if err != nil {
		return nil, err
	}

	// if err := Client.Schema.Create(
	// 	context.Background(),
	// 	migrate.WithDropIndex(true),
	// 	migrate.WithDropColumn(true),
	// ); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Client: client}}))
	return srv, nil
}

func Close(client *ent.Client) {
	if client != nil {
		client.Close()
	}
}
