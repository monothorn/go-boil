package posts

import (
	"context"

	"monothorn/go-boil/posts/models"
)

type PostsRepository interface {
	GetPosts(ctx context.Context) ([]*models.Post, error)
}
