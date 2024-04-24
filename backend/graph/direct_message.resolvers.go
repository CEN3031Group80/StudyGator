package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"study-gator-backend/graph/model"
)

// ID is the resolver for the id field.
func (r *directMessageResolver) ID(ctx context.Context, obj *model.DirectMessage) (string, error) {
	return obj.GetID(), nil
}

// Members is the resolver for the members field.
func (r *directMessageResolver) Members(ctx context.Context, obj *model.DirectMessage) ([]*model.User, error) {
	var users []model.User

	tx := model.DB.Joins("DirectMessageMember", model.DB.Where(&model.DirectMessageMember{DirectMessageID: int(obj.ID)})).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var userPointers []*model.User
	for _, user := range users {
		userPointers = append(userPointers, &user)
	}

	return userPointers, nil
}

// Posts is the resolver for the posts field.
func (r *directMessageResolver) Posts(ctx context.Context, obj *model.DirectMessage, limit int, offset int) ([]*model.DirectMessagePost, error) {
	var posts []model.DirectMessagePost

	tx := model.DB.Limit(limit).Offset(offset).Find(&posts, "direct_message_id = ?", obj.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var postPointers []*model.DirectMessagePost
	for _, post := range posts {
		postPointers = append(postPointers, &post)
	}

	return postPointers, nil
}

// ID is the resolver for the id field.
func (r *directMessagePostResolver) ID(ctx context.Context, obj *model.DirectMessagePost) (string, error) {
	return obj.GetID(), nil
}

// Sender is the resolver for the sender field.
func (r *directMessagePostResolver) Sender(ctx context.Context, obj *model.DirectMessagePost) (*model.User, error) {
	var user model.User
	tx := model.DB.First(&user, obj.UserID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

// DirectMessage returns DirectMessageResolver implementation.
func (r *Resolver) DirectMessage() DirectMessageResolver { return &directMessageResolver{r} }

// DirectMessagePost returns DirectMessagePostResolver implementation.
func (r *Resolver) DirectMessagePost() DirectMessagePostResolver {
	return &directMessagePostResolver{r}
}

type directMessageResolver struct{ *Resolver }
type directMessagePostResolver struct{ *Resolver }
