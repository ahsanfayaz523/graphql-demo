package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/ahsanfayaz523/graphql-demo/db/mongo"

	"github.com/ahsanfayaz523/graphql-demo/graph/generated"
	"github.com/ahsanfayaz523/graphql-demo/graph/model"
)

var db = mongo.Connect()

// CreateAnimal is the resolver for the createAnimal field.
func (r *mutationResolver) CreateAnimal(ctx context.Context, input *model.NewAnimal) (*model.Animals, error) {
	return db.Save(input), nil
}

// SingleAnimal is the resolver for the SingleAnimal field.
func (r *queryResolver) SingleAnimal(ctx context.Context, id string) (*model.Animals, error) {
	return db.FetchAnimal(id), nil
}

// AllAnimals is the resolver for the AllAnimals field.
func (r *queryResolver) AllAnimals(ctx context.Context) ([]*model.Animals, error) {
	return db.FetchAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
