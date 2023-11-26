package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	// "strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/fadedreams/xclone"
	"github.com/fadedreams/xclone/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	res, err := r.AuthService.Register(ctx, xclone.RegisterInput{
		Email:           input.Email,
		Username:        input.Username,
		Password:        input.Password,
		ConfirmPassword: input.ConfirmPassword,
	})
	if err != nil {
		switch {
		case errors.Is(err, xclone.ErrValidation) ||
			errors.Is(err, xclone.ErrEmailTaken) ||
			errors.Is(err, xclone.ErrUsernameTaken):
			return nil, buildBadRequestError(ctx, err)
		default:
			return nil, err
		}
	}

	return mapAuthResponse(res), nil
	panic(fmt.Errorf("not implemented: Register - register"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// CreateTweet is the resolver for the createTweet field.
func (r *mutationResolver) CreateTweet(ctx context.Context, input model.CreateTweetInput) (*model.Tweet, error) {
	panic(fmt.Errorf("not implemented: CreateTweet - createTweet"))
}

// CreateReply is the resolver for the createReply field.
func (r *mutationResolver) CreateReply(ctx context.Context, parentID string, input model.CreateTweetInput) (*model.Tweet, error) {
	panic(fmt.Errorf("not implemented: CreateReply - createReply"))
}

// DeleteTweet is the resolver for the deleteTweet field.
func (r *mutationResolver) DeleteTweet(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteTweet - deleteTweet"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// Tweets is the resolver for the tweets field.
func (r *queryResolver) Tweets(ctx context.Context) ([]*model.Tweet, error) {
	panic(fmt.Errorf("not implemented: Tweets - tweets"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.

// func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
// 	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
// }
// func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
// 	panic(fmt.Errorf("not implemented: Todos - todos"))
// }

func buildBadRequestError(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": http.StatusBadRequest,
		},
	}
}

func mapAuthResponse(a xclone.AuthResponse) *model.AuthResponse {
	return &model.AuthResponse{
		AccessToken: a.AccessToken,
		User:        mapUser(a.User),
	}
}

func mapUser(u xclone.User) *model.User {
	return &model.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt, // Commented out for testing
	}

}
