package posts

import (
	"context"
	"monothorn/go-boil/posts/models"
)

type PostsUseCase interface {
	GetPosts(ctx context.Context) ([]*models.Post, error)
}
