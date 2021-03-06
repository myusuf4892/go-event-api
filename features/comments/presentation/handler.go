package presentation

import (
	"net/http"
	"project3/eventapp/features/comments"
	_request_comment "project3/eventapp/features/comments/presentation/request"
	_response_comment "project3/eventapp/features/comments/presentation/response"
	"strconv"

	"project3/eventapp/helper"
	"project3/eventapp/middlewares"

	// "strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentBusiness comments.Business
}

func NewCommentHandler(business comments.Business) *CommentHandler {
	return &CommentHandler{
		commentBusiness: business,
	}
}

func (h *CommentHandler) Add(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get user id"))
	}

	comment := _request_comment.Comment{}
	err_bind := c.Bind(&comment)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind insert data"))
	}

	commentCore := _request_comment.ToCore(comment)
	commentCore.UserID = userID_token

	row, err := h.commentBusiness.AddComment(commentCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed insert your comment"))
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed insert your comment"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success insert your comment"))
}

func (h *CommentHandler) Get(c echo.Context) error {
	eventId, _ := strconv.Atoi(c.Param("id"))
	limit := 5
	offset, _ := strconv.Atoi(c.QueryParam("page"))

	result, total, err := h.commentBusiness.GetCommentByIdEvent(limit, offset, eventId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed get all comment events"))
	}
	respons := _response_comment.FromCoreList(result)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithDataPage("success get all comment events", total, respons))
}
