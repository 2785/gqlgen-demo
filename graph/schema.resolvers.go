package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/2785/gqlgen-demo/graph/generated"
	"github.com/2785/gqlgen-demo/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) AddUser(ctx context.Context, addUser *model.AddUser) (bool, error) {
	// this should be taken from a service wrapping a repository rather than everything in resolver
	// but I can't be arsed
	for _, v := range r.UsersStore {
		if v.Name == *addUser.Name {
			return false, fmt.Errorf("cannot add: user %s exists", *addUser.Name)
		}
	}

	if *addUser.Age > 80 || *addUser.Age < 18 {
		return false, fmt.Errorf("only serve users between 18 and 80 years old, cannot add user with age %v", addUser.Age)
	}

	r.UsersStore = append(r.UsersStore, &model.User{ID: uuid.New().String(), Age: *addUser.Age, Name: *addUser.Name})
	return true, nil
}

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	for _, v := range r.UsersStore {
		if v.ID == *id {
			return v, nil
		}
	}
	return nil, fmt.Errorf("user '%s' not found", *id)
}

func (r *queryResolver) Users(ctx context.Context, search *model.UserSearchParam) ([]*model.User, error) {
	return r.UsersStore, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
