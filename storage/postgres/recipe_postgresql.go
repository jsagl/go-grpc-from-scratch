package storage

import (
	"context"
	"database/sql"
	"fmt"
	v1 "github.com/jsagl/go-grpc-from-scratch/pkg/api/v1"
	"github.com/jsagl/go-grpc-from-scratch/storage"
)

type PostgresRecipeStore struct {
	Connection *sql.DB
}

func NewPostgresRecipeStore(connection *sql.DB) storage.RecipeStore {
	return &PostgresRecipeStore{Connection: connection}
}

func (store *PostgresRecipeStore) Read(ctx context.Context,recipeId int64) (*v1.Recipe, error) {
	var recipe v1.Recipe

	query := `SELECT id, title, description FROM recipes WHERE ID = $1`
	if err := store.Connection.QueryRowContext(ctx, query, recipeId).Scan(&recipe.Id, &recipe.Title, &recipe.Description); err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (store *PostgresRecipeStore) ReadAll(ctx context.Context) ([]*v1.Recipe, error) {
	query:= "SELECT id, title, description FROM recipes"
	rows, err := store.Connection.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	recipes := make([]*v1.Recipe, 0)

	for rows.Next() {
		var recipe v1.Recipe
		err := rows.Scan(&recipe.Id, &recipe.Title, &recipe.Description)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}