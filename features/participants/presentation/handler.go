package presentation

import (
	"fmt"
	"net/http"
	"project3/eventapp/features/participants"
	_request_participant "project3/eventapp/features/participants/presentation/request"
	_response_participant "project3/eventapp/features/participants/presentation/response"
	"project3/eventapp/helper"
	"project3/eventapp/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ParticipantHandler struct {
	participantBusiness participants.Business
}

func NewParticipantHandler(business participants.Business) *ParticipantHandler {
	return &ParticipantHandler{
		participantBusiness: business,
	}
}

func (h *ParticipantHandler) Joined(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed insert data"))
	}

	participant := _request_participant.Participant{}
	err_bind := c.Bind(&participant)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("success insert participant"))
	}

	participantCore := _request_participant.ToCore(participant)
	participantCore.UserID = userID_token

	err := h.participantBusiness.AddParticipant(participantCore)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed join"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success join"))

}

func (h *ParticipantHandler) GetAllEventParticipant(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed insert data"))
	}

	result, err := h.participantBusiness.GetAllEventbyParticipant(userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed get all your event"))
	}

	response := _response_participant.FromCoreList(result)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success get all your events", response))

}

func (h *ParticipantHandler) DeleteEventbyParticipant(c echo.Context) error {
	idParticipant, _ := strconv.Atoi(c.Param("id"))

	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed get user id"))
	}

	result := h.participantBusiness.DeleteParticipant(idParticipant, userID_token)
	if result != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to delete your event"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success delete to delete your event"))

}
