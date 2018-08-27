package usecase

import (
	"context"
	"time"

	"monothorn/go-boil/posts"

	"monothorn/go-boil/posts/models"

	"monothorn/go-boil/posts/repository"
)

type postsUseCase struct {
	cmRepos        posts.PostsRepository
	contextTimeout time.Duration
}

func NewPostsUseCase(cm repository.PostsRepository, timeout time.Duration) repository.PostsRepository {
	return &postsUseCase{
		cm, timeout,
	}
}

func (cm *postsUseCase) GetPosts(ctx context.Context) ([]*models.Post, error) {
	listP, err := cm.cmRepos.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	return listP, nil
}
