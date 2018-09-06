package repository

import (
	"context"

	"monothorn/go-boil/posts/models"
	"monothorn/go-boil/utilities"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type PostsRepository interface {
	GetPosts(ctx context.Context) ([]*models.Post, error)
	AddPost(ctx context.Context, postData *models.Post) (string, error)
}

type myPostsRepository struct {
	source string
}

func NewPostsRepository(source string) PostsRepository {
	return &myPostsRepository{"apiv1"}
}

func (cm *myPostsRepository) GetPosts(ctx context.Context) ([]*models.Post, error) {
	db, _ := utilities.GetInstance(viper.GetString("sql.dsn"))
	rows, err := db.Query("SELECT postTitle,postBody FROM test_db.posts")
	result := make([]*models.Post, 0)
	for rows.Next() {
		p := &models.Post{}
		err = rows.Scan(&p.PostTitle, &p.PostBody)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, p)
	}

	return result, nil
}

func (cm *myPostsRepository) AddPost(ctx context.Context, postData *models.Post) (string, error) {
	db, _ := utilities.GetInstance(viper.GetString("sql.dsn"))
	stmt, err := db.Prepare("INSERT into test_db.posts(postTitle,postBody) VALUES(?,?)")
	if err != nil {
		logrus.Error(err)
		return "failure!", nil
	}
	defer stmt.Close()

	_, errr := stmt.Exec(postData.PostTitle, postData.PostBody)
	if errr != nil {
		logrus.Error(errr)
		return "failure!", nil
	}

	return "success!", nil

}
