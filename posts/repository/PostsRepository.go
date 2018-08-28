package repository

import (
	"context"

	"monothorn/go-boil/posts/models"
)

type PostsRepository interface {
	GetPosts(ctx context.Context) ([]*models.Post, error)
}

type myPostsRepository struct {
	source string
}

func NewPostsRepository(source string) PostsRepository {
	return &myPostsRepository{"apiv1"}
}

func (cm *myPostsRepository) GetPosts(ctx context.Context) ([]*models.Post, error) {
	result := make([]*models.Post, 0)
	t := new(models.Post)
	t.PostTitle = "Big Data"
	t.PostBody = "Very good!!"
	result = append(result, t)

	return result, nil
}
