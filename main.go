package main

import (
	"context"
	"database/sql"
	"github.com/jsagl/go-grpc-from-scratch/api/service/v1"
	"github.com/jsagl/go-grpc-from-scratch/server/grpc"
	"github.com/jsagl/go-grpc-from-scratch/server/rest"
	"github.com/jsagl/go-grpc-from-scratch/storage/postgres"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("postgres", "postgres://postgres:@localhost:5432/go_recipes")
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	recipeStore := storage.NewPostgresRecipeStore(db)
	v1RecipeServiceServer := v1.NewRecipeServiceServer(recipeStore)

	// run HTTP gateway
	go func() {
		_ = rest.StartHTTP(ctx)
	}()

	grpc.StartGRPC(ctx, v1RecipeServiceServer, "8080")

}