//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ext, err := entgql.NewExtension(
		entgql.WithSchemaPath("../graphql/graph/ent.graphql"),
		entgql.WithSchemaGenerator(),
		entgql.WithConfigPath("../graphql/gqlgen.yml"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ext))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
