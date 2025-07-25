package handler

import (
	domain "gomibakokun_backend/domain/trashcan"
	interfaces "gomibakokun_backend/interfaces"
	"gomibakokun_backend/usecase"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

type TrashcanHandler interface {
	HandleTrashcanCreate(c echo.Context) error
	HandleTrashcansInRange(c echo.Context) error
	HandleTrashcanDelete(c echo.Context) error
}

type trashcanHandler struct {
	trashcanUsecase usecase.TrashcanUseCase
}

func NewTrashcanHandler(tu usecase.TrashcanUseCase) TrashcanHandler {
	return &trashcanHandler{
		trashcanUsecase: tu,
	}
}

func (th trashcanHandler) HandleTrashcanCreate(c echo.Context) error {
	var createTrashcanRequest interfaces.CreateTrashcanRequest
	if err := c.Bind(&createTrashcanRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "Invalid input"})
	}

	ctx := c.Request().Context()
	trashcanInput := &usecase.TrashcanInput{
		Latitude:        createTrashcanRequest.Latitude,
		Longitude:       createTrashcanRequest.Longitude,
		Image:           createTrashcanRequest.Image,
		TrashType:       createTrashcanRequest.TrashType,
		NearestBuilding: createTrashcanRequest.NearestBuilding,
		Note:            createTrashcanRequest.Note,
		SelectedButton:  createTrashcanRequest.SelectedButton,
	}

	err := th.trashcanUsecase.CreateTrashcan(ctx, trashcanInput)
	if err != nil {
		if err == domain.ErrInvalidInput {
			return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "Invalid input"})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{"success": false, "message": "Failed to create trashcan"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"success": true})
}

func (th trashcanHandler) HandleTrashcansInRange(c echo.Context) error {
	latitude := c.QueryParam("latitude")
	longitude := c.QueryParam("longitude")

	ctx := c.Request().Context()

	range_radius := 20000 //TODO:リリース時治す

	latitudeFloat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "Invalid latitude"})
	}

	longitudeFloat, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"success": false, "message": "Invalid longitude"})
	}

	trashcanOutputs, err := th.trashcanUsecase.GetTrashcansInRange(ctx, latitudeFloat, longitudeFloat, float64(range_radius))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"success": false})
	}

	trashcanResponses := make([]interfaces.TrashcanResponse, len(trashcanOutputs))
	for i, trashcanOutput := range trashcanOutputs {
		trashcanResponses[i] = interfaces.TrashcanResponse{
			ID:              trashcanOutput.ID,
			Latitude:        trashcanOutput.Latitude,
			Longitude:       trashcanOutput.Longitude,
			Image:           trashcanOutput.Image,
			TrashType:       trashcanOutput.TrashType,
			NearestBuilding: trashcanOutput.NearestBuilding,
			Note:            trashcanOutput.Note,
			SelectedButton:  trashcanOutput.SelectedButton,
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true, "trashcans": trashcanResponses})
}

func (th trashcanHandler) HandleTrashcanDelete(c echo.Context) error {
	id := c.QueryParam("id")

	ctx := c.Request().Context()

	err := th.trashcanUsecase.DeleteTrashcan(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"success": false})
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}
