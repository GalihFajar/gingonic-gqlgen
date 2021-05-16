package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql_todos/Config"
	"graphql_todos/graph/generated"
	"graphql_todos/graph/model"

	_ "github.com/go-sql-driver/mysql"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	user  := model.User{
		Name: input.Name,
		Address: input.Address,
		Email: input.Email,
		Phone: input.Phone,
	}
	err := Config.DB.Create(&user).Error
	return &user, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var user []*model.User
	var temp []model.User
	err := Config.DB.Find(&temp).Error
	for i := range temp{
		tempPointer := &temp[i]
		user = append(user, tempPointer)
	}
	return user, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
