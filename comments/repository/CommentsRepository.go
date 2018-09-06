package repository

import (
	"context"
	"monothorn/go-boil/comments/models"
	"monothorn/go-boil/utilities"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type CommentsRepository interface {
	GetCommentsByPost(ctx context.Context, postID int) ([]*models.Comments, error)
	AddCommentToPost(ctx context.Context, commentData *models.Comments) (string, error)
}

type myCommentsRepository struct {
	source string
}

func NewCommentsRepository(source string) CommentsRepository {
	return &myCommentsRepository{"apiv1"}
}

func (cmR *myCommentsRepository) GetCommentsByPost(ctx context.Context, postId int) ([]*models.Comments, error) {
	db, _ := utilities.GetInstance(viper.GetString("sql.dsn"))
	result := make([]*models.Comments, 0)
	stmt, err := db.Prepare("SELECT postId,postCommentTitle,postCommentBody FROM test_db.comments WHERE postId = ?")
	defer stmt.Close()

	rows, err := stmt.Query(postId)

	defer rows.Close()

	for rows.Next() {
		m := &models.Comments{}
		err = rows.Scan(&m.PostID, &m.PostCommentTitle, &m.PostCommentBody)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, m)
	}
	return result, nil
}

func (cmR *myCommentsRepository) AddCommentToPost(ctx context.Context, commentData *models.Comments) (string, error) {
	db, _ := utilities.GetInstance(viper.GetString("sql.dsn"))
	stmt, err := db.Prepare("INSERT into test_db.comments(postId,postCommentTitle,postCommentBody) VALUES (?,?,?)")
	if err != nil {
		logrus.Error(err)
		return "failure!", nil
	}
	defer db.Close()
	defer stmt.Close()

	_, errr := stmt.Query(commentData.PostID, commentData.PostCommentTitle, commentData.PostCommentBody)
	if errr != nil {
		logrus.Error(errr)
		return "failure!", nil
	}

	return "success!", nil
}
