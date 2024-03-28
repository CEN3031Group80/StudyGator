package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"study-gator-backend/graph/gqlcontext"
	"study-gator-backend/graph/model"
)

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user := gqlcontext.UserFromContext(ctx)
	return &model.User{
		ID:        user.ID,
		AvatarURL: user.Picture,
		AuthInfo: &model.AuthInfo{
			Provider: model.AuthProvidersGithub,
			Name:     user.Name,
			Email:    user.Email,
		},
		Profile: &model.Profile{
			FirstName:      "Test",
			LastName:       "Test",
			School:         "UF",
			GraduationYear: 2025,
		},
	}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
