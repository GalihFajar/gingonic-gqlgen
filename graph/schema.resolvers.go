package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql_todos/Config"
	"graphql_todos/graph/generated"
	"graphql_todos/graph/model"
	"graphql_todos/utils"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	var user model.User
	user.Name = input.Name
	user.Address = input.Address
	user.Email = input.Email
	user.Phone = input.Phone
	err := Config.DB.Create(&user).Error
	returnedUser := &user
	return returnedUser, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var user []*model.User
	err := Config.DB.Find(user).Error
	gc, cerr := utils.GinContextFromContext(ctx)
	if err != nil || cerr != nil {
		fmt.Println(err.Error())
		gc.AbortWithStatus(http.StatusNotFound)
	} else {
		gc.JSON(http.StatusOK, user)
	}
	return user, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
