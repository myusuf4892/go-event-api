package presentation

import (
	"net/http"
	"project3/eventapp/features/auth"
	"project3/eventapp/features/auth/presentation/request"
	"project3/eventapp/features/auth/presentation/response"
	"project3/eventapp/helper"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userBusiness auth.Business
}

func NewAuthHandler(business auth.Business) *AuthHandler {
	return &AuthHandler{
		userBusiness: business,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	reqBody := request.User{}
	errBind := c.Bind(&reqBody)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get bind data",
		})
	}

	authCore := request.ToCore(reqBody)
	token, id, err := h.userBusiness.Login(authCore)

	result := response.ToResponse(id, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get token data" + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("login success", result))
}
