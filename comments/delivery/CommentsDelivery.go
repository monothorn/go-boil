package delivery

import (
	"monothorn/go-boil/comments/models"
	"monothorn/go-boil/comments/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:message`
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

type CommentsHandler struct {
	cmUC usecase.CommentsUseCase
}

func (cmUC *CommentsHandler) GetComments(c echo.Context) error {
	ctx := c.Request().Context()
	postID, _ := strconv.Atoi(c.Param("postID"))
	listAr, err := cmUC.cmUC.GetCommentsByPost(ctx, postID)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listAr)
}

func (cmUC *CommentsHandler) AddComments(c echo.Context) error {
	ctx := c.Request().Context()
	commentData := models.Comments{}
	if err := c.Bind(&commentData); err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	postID, _ := strconv.Atoi(c.Param("postID"))
	commentData.PostID = postID

	msg, err := cmUC.cmUC.AddCommentToPost(ctx, &commentData)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, msg)
}

func NewCommentsHandler(r *echo.Echo, cmUC usecase.CommentsUseCase) {
	handler := &CommentsHandler{cmUC: cmUC}
	r.GET("/post/:postID", handler.GetComments)
	r.POST("/post/:postID/comment", handler.AddComments)
}
