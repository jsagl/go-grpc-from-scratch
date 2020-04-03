package v1

import (
	"context"
	"github.com/jsagl/go-grpc-from-scratch/api/proto/v1"
	"github.com/jsagl/go-grpc-from-scratch/storage"
)

type recipeServiceServer struct {
	store storage.RecipeStore
}

func NewRecipeServiceServer(store storage.RecipeStore) v1.RecipeServiceServer {
	return &recipeServiceServer{store: store}
}

func (server *recipeServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	recipe, err := server.store.Read(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	response := &v1.ReadResponse{Recipe: recipe}

	return response, nil
}

func (server *recipeServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	recipes, err := server.store.ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	response := &v1.ReadAllResponse{Recipes: recipes}

	return response, nil
}