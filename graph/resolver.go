package graph

import "github.com/2785/gqlgen-demo/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UsersStore []*model.User
}
