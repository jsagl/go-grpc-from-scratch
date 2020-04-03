package	storage

import (
	"context"
	"github.com/jsagl/go-grpc-from-scratch/api/proto/v1"
)

type RecipeStore interface {
	Read(ctx context.Context, id int64) (*v1.Recipe, error)
	ReadAll(ctx context.Context) ([]*v1.Recipe, error)
}