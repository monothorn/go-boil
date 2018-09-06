package usecase

import (
	"context"
	"time"

	"monothorn/go-boil/posts/models"

	"monothorn/go-boil/posts/repository"
)

type PostsUseCase interface {
	GetPosts(ctx context.Context) ([]*models.Post, error)
	AddPost(ctx context.Context, postData *models.Post) (string, error)
}

type postsUseCase struct {
	cmRepos        repository.PostsRepository
	contextTimeout time.Duration
}

func NewPostsUseCase(cm repository.PostsRepository, timeout time.Duration) PostsUseCase {
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

func (cm *postsUseCase) AddPost(ctx context.Context, postData *models.Post) (string, error) {
	msg, err := cm.cmRepos.AddPost(ctx, postData)
	if err != nil {
		return "failure!", err
	}
	return msg, nil
}
