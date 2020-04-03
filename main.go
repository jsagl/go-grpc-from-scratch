package main

import (
	"context"
	"github.com/jsagl/go-grpc-from-scratch/pkg/protocol/grpc"
	"github.com/jsagl/go-grpc-from-scratch/pkg/service/v1"
	"github.com/jsagl/go-grpc-from-scratch/storage/postgres"
	"log"
)

func main() {
	ctx := context.Background()
	db, err := storage.NewPostgresConnection()
	if err != nil {
		log.Panic(err)
	}

	recipeStore := storage.NewPostgresRecipeStore(db)
	v1RecipeServiceServer := v1.NewRecipeServiceServer(recipeStore)

	grpc.RunServer(ctx, v1RecipeServiceServer, "8080")
}