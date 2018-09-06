package usecase

import (
	"context"
	"monothorn/go-boil/comments/models"
	"monothorn/go-boil/comments/repository"
	"time"
)

type CommentsUseCase interface {
	GetCommentsByPost(ctx context.Context, postID int) ([]*models.Comments, error)
	AddCommentToPost(ctx context.Context, commentData *models.Comments) (string, error)
}

type commentsUseCase struct {
	cmRRepo        repository.CommentsRepository
	contextTimeout time.Duration
}

func NewCommentsUseCase(cmRRepo repository.CommentsRepository, timeout time.Duration) CommentsUseCase {
	return &commentsUseCase{
		cmRRepo,
		timeout,
	}
}

func (cmUC *commentsUseCase) GetCommentsByPost(ctx context.Context, postID int) ([]*models.Comments, error) {
	listC, err := cmUC.cmRRepo.GetCommentsByPost(ctx, postID)
	if err != nil {
		return nil, err
	}
	return listC, err
}

func (cmUC *commentsUseCase) AddCommentToPost(ctx context.Context, commentData *models.Comments) (string, error) {
	msg, err := cmUC.cmRRepo.AddCommentToPost(ctx, commentData)
	if err != nil {
		return "failure!", err
	}
	return msg, nil
}
