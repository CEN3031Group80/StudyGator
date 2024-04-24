package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"study-gator-backend/graph/model"
)

// ID is the resolver for the id field.
func (r *postResolver) ID(ctx context.Context, obj *model.Post) (string, error) {
	return obj.GetID(), nil
}

// DateCreated is the resolver for the dateCreated field.
func (r *postResolver) DateCreated(ctx context.Context, obj *model.Post) (string, error) {
	return obj.CreatedAt.String(), nil
}

// DateUpdated is the resolver for the dateUpdated field.
func (r *postResolver) DateUpdated(ctx context.Context, obj *model.Post) (string, error) {
	return obj.UpdatedAt.String(), nil
}

// Attachments is the resolver for the attachments field.
func (r *postResolver) Attachments(ctx context.Context, obj *model.Post) ([]*model.PostAttachment, error) {
	var attachments []model.PostAttachment

	tx := model.DB.Find(&attachments, "post_id = ?", obj.ID)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var postAttachmentPointers []*model.PostAttachment
	for _, attachment := range attachments {
		postAttachmentPointers = append(postAttachmentPointers, &attachment)
	}

	return postAttachmentPointers, nil
}

// ID is the resolver for the id field.
func (r *postAttachmentResolver) ID(ctx context.Context, obj *model.PostAttachment) (string, error) {
	return obj.GetID(), nil
}

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// PostAttachment returns PostAttachmentResolver implementation.
func (r *Resolver) PostAttachment() PostAttachmentResolver { return &postAttachmentResolver{r} }

type postResolver struct{ *Resolver }
type postAttachmentResolver struct{ *Resolver }
