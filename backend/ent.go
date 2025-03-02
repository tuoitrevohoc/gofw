//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithSchemaPath("schema/graphql/ent.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(ex),
		entc.FeatureNames("schema/migrate"),
	}

	if err := entc.Generate("./schema/models", &gen.Config{
		Target:  "./gen/go/ent",
		Package: "github.com/tuoitrevohoc/gofw/backend/gen/go/ent",
	}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
