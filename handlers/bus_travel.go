package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lleonesouza/bus-travel/modules/dtos"
	er "github.com/lleonesouza/bus-travel/modules/errors"
	"github.com/lleonesouza/bus-travel/prisma/db"
	"github.com/lleonesouza/bus-travel/services"
)

type BusTravelHandler struct {
	service *services.BusTravelService
}

func (h *BusTravelHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/:id", h.Get)
	g.GET("", h.GetAll)
	g.POST("", h.Create)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

// @Description	Find bus-travel.
// @Tags			bus-travel
// @Accept			json
// @Produce		json
// @Param id path int true "BusTravel ID"
// @Success		200	{object}	dtos.BusTravel
// @Failure		400	{object}	er.BadRequestError
// @Failure		404	{object}	er.NotFoundError
// @Failure		500	{object}	er.InternalServerError
// @Router			/bus-travels/{id} [get]
func (h *BusTravelHandler) Get(c echo.Context) error {
	// Get the `id` parameter from the URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, er.BadRequest("Invalid id parameter"))
	}

	// Retrieve the bus travel resource from the database
	busTravel, err := h.service.Get(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return c.JSON(http.StatusNotFound, er.NotFound("BusTravel", string(id)))
		}
		return c.JSON(http.StatusInternalServerError, er.InternalServer(err.Error()))
	}

	return c.JSON(http.StatusOK, busTravel)
}

// @Description	Create a bus-travel.
// @Tags			bus-travel
// @Accept			json
// @Produce		json
// @Param			body	body		dtos.BusTravelCreateInput	true	"Create BusTravel Input"
// @Success		201		{object}	dtos.BusTravel
// @Failure		400		{object}	er.BadRequestError
// @Failure		500		{object}	er.InternalServerError
// @Router			/bus-travels [post]
func (h *BusTravelHandler) Create(c echo.Context) error {
	// Parse input data from request body
	var input dtos.BusTravelCreateInput
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, er.BadRequest("Invalid input data"))
	}

	// Validate input data
	err = input.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, er.BadRequest(err.Error()))
	}

	// Create new bus travel resource in DB
	createdBusTravel, err := h.service.Create(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, er.InternalServer(err.Error()))
	}

	return c.JSON(http.StatusCreated, createdBusTravel)
}

// @Description	Update BusTravel
// @Tags			bus-travel
// @Accept			json
// @Produce		json
// @Param id path int true "BusTravel ID"
// @Param			body	body		dtos.BusTravelUpdateInput	true	"BusTravel"
// @Success		200		{object}	dtos.BusTravel
// @Failure		400		{object}	er.BadRequestError
// @Failure		404		{object}	er.NotFoundError
// @Failure		500		{object}	er.InternalServerError
// @Router			/bus-travels/{id} [put]
func (h *BusTravelHandler) Update(c echo.Context) error {
	// Get the bus travel ID from the URL parameter
	busTravelID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, er.BadRequest("Invalid bus travel ID"))
	}

	// Parse input data from request body
	var input dtos.BusTravelUpdateInput
	err = c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, er.BadRequest("Invalid input data"))
	}

	// Validate input data
	err = input.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, er.BadRequest(err.Error()))
	}

	// Update the bus travel resource in DB
	updatedBusTravel, err := h.service.Update(c.Request().Context(), busTravelID, input)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return c.JSON(http.StatusNotFound, er.NotFound("BusTravel", string(busTravelID)))
		}
		return c.JSON(http.StatusInternalServerError, er.InternalServer(err.Error()))
	}

	return c.JSON(http.StatusOK, updatedBusTravel)
}

// @Description	Delete BusTravel
// @Tags			bus-travel
// @Accept			json
// @Produce			json
// @Param id path int true "BusTravel ID"
// @Success 204 "No content"
// @Failure			400		object	er.BadRequestError
// @Failure			404		object	er.NotFoundError
// @Failure			500		object	er.InternalServerError
// @Router			/bus-travels/{id} [delete]
func (h *BusTravelHandler) Delete(c echo.Context) error {
	// Get the ID of the bus travel to be deleted
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, er.BadRequest("Invalid ID"))
	}

	// Delete the bus travel resource from DB
	err = h.service.Delete(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return c.JSON(http.StatusNotFound, er.NotFound("BusTravel", string(id)))
		}
		return c.JSON(http.StatusInternalServerError, er.InternalServer(err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}

// @Description	Get all bus travels
// @Tags			bus-travel
// @Accept			json
// @Produce		json
// @Success		200		{array}		dtos.BusTravel
// @Failure		400		{object}	er.BadRequestError
// @Failure		500		{object}	er.InternalServerError
// @Router			/bus-travels [get]
func (h *BusTravelHandler) GetAll(c echo.Context) error {
	// Get all bus travels from the service
	busTravels, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, er.InternalServer(err.Error()))
	}

	// Return the bus travels as response
	return c.JSON(http.StatusOK, busTravels)
}
