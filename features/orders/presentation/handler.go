package presentation

import (
	"log"
	"net/http"
	"project/e-comerce/features/orders"
	"project/e-comerce/features/orders/presentation/request"
	"project/e-comerce/helper"
	"project/e-comerce/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderBusiness orders.Bussiness
}

func NewOrderHandler(business orders.Bussiness) *OrderHandler {
	return &OrderHandler{
		orderBusiness: business,
	}
}

func (h OrderHandler) AddOrder(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get id user"))
	}

	reqData := request.Order{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind data"))
	}

	log.Print("request = ",reqData.Address)

	dataCore := request.ToCore(reqData)
	dataCore.UserID = userID_token

	log.Print("core address = ", dataCore.Address)
	err := h.orderBusiness.AddOrder(dataCore, reqData.CartID, userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success"))
}

func (h OrderHandler) GetOrder(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get user id"))
	}
	dataOrder, err := h.orderBusiness.GetOrder(userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}

	//responOrder := response.FromCoreList(dataOrder)

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success", dataOrder))
}

func (h OrderHandler) GeOrderDetailByOrderID(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get user id"))
	}

	idOrder, errParam := strconv.Atoi(c.Param("orderid"))
	if errParam != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get param order id"))
	}

	dataOrderDetail, err := h.orderBusiness.GetOrderDetail(idOrder)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success", dataOrderDetail))
}

func (h OrderHandler) ConfirmOrder(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get user id"))
	}

	orderId, errParam := strconv.Atoi(c.Param("orderid"))
	if errParam != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get param order id"))
	}
	payReq := request.Payment{}
	errBind := c.Bind(&payReq)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get bind payment data"))
	}

	payCore := request.ToPaymentCore(payReq)

	errRespon := h.orderBusiness.ConfirmOrder(payCore, orderId, userID_token)
	if errRespon != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errRespon.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success"))
}
func (h OrderHandler) CancelOrder(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get user id"))
	}

	orderId, errParam := strconv.Atoi(c.Param("orderid"))
	if errParam != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get param order id"))
	}

	errRespon := h.orderBusiness.CancelOrder(orderId, userID_token)
	if errRespon != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errRespon.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success"))
}
