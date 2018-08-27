package delivery

import (
	cmUseCase "monothorn/go-boil/posts"
	"net/http"

	"github.com/bxcodec/go-clean-arch/models"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:

		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

type PostsHandler struct {
	cmUC cmUseCase.PostsUseCase
}

func (cm *PostsHandler) GetPosts(c echo.Context) error {
	ctx := c.Request().Context()
	listAr, err := cm.cmUC.GetPosts(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listAr)
}

// NewCareerMovesHandler handles routing for feedback module
func NewPostsHandler(r *echo.Echo, cmUC cmUseCase.PostsUseCase) {
	handler := &PostsHandler{
		cmUC: cmUC,
	}
	r.GET("/", handler.GetPosts)
}
