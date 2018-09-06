package delivery

import (
	"monothorn/go-boil/posts/models"
	"monothorn/go-boil/posts/usecase"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type PostData struct {
	postTitle string
	postBody  string
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

type PostsHandler struct {
	cmUC usecase.PostsUseCase
}

func (cm *PostsHandler) GetPosts(c echo.Context) error {
	ctx := c.Request().Context()
	listAr, err := cm.cmUC.GetPosts(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listAr)
}

func (cm *PostsHandler) AddPost(c echo.Context) error {
	ctx := c.Request().Context()
	postData := models.Post{}
	if err := c.Bind(&postData); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	msg, err := cm.cmUC.AddPost(ctx, &postData)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, msg)
}

// NewPostsHandler handles routing for feedback module
func NewPostsHandler(r *echo.Echo, cmUC usecase.PostsUseCase) {
	handler := &PostsHandler{
		cmUC: cmUC,
	}
	r.GET("/posts", handler.GetPosts)
	r.POST("/post", handler.AddPost)
}
