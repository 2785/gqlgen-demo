package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/2785/gqlgen-demo/graph/generated"
	"github.com/2785/gqlgen-demo/graph/model"
)

func (r *mutationResolver) AddBuyer(ctx context.Context, buyer *model.AddBuyer) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Account(ctx context.Context, id *int) (model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchAccount(ctx context.Context, params *model.AccountSearchParam) ([]model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
