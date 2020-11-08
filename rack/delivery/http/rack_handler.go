package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/zarszz/warehouse-rest-api/constant"
	"github.com/zarszz/warehouse-rest-api/domain"
	"github.com/zarszz/warehouse-rest-api/utils"
)

// RackHandler  represent the httphandler for article
type RackHandler struct {
	RackUsecase domain.RackUseCase
}

// NewRackHandler will initialize the racks/ resources endpoint
func NewRackHandler(e *echo.Echo, rackUsecase domain.RackUseCase) {
	handler := &RackHandler{
		RackUsecase: rackUsecase,
	}
	e.GET("/racks", handler.FetchRack)
	e.POST("/racks", handler.Store)
	e.GET("/racks/:id", handler.GetByID)
	e.DELETE("/racks/:id", handler.Delete)
	e.PUT("/racks/:id", handler.Update)
}

// FetchCategory will fetch the racks based on given params
func (a *RackHandler) FetchRack(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()

	listRacks, nextCursor, err := a.RackUsecase.Fetch(ctx, cursor, int64(num))
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.CATEGORY_GET_FAILED, utils.GetStatusCode(err))
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, listRacks)
}

// GetByID will get rack by given id
func (a *RackHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	category, err := a.RackUsecase.GetByID(ctx, id)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), utils.GetStatusCode(err))
	}

	return utils.HandleResponseGet(c, constant.SUCCESS_LOAD_DATA, constant.SUCCESS, http.StatusOK, category)
}

func isRequestValid(m *domain.Rack) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the category by given request body
func (a *RackHandler) Store(c echo.Context) (err error) {
	var rack domain.Rack
	err = c.Bind(&rack)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), http.StatusUnprocessableEntity)
	}

	var ok bool
	if ok, err = isRequestValid(&rack); !ok {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), http.StatusBadRequest)
	}

	rack.CreatedAt = time.Now()
	rack.UpdatedAt = time.Now()
	rack.ID = utils.GenerateSHA1(rack.Name, time.Now().String())

	ctx := c.Request().Context()
	err = a.RackUsecase.Store(ctx, &rack)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), utils.GetStatusCode(err))
	}

	return c.JSON(http.StatusCreated, rack)
}

// Update will update the rack given with specific id
func (a *RackHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var rack domain.Rack
	if err := c.Bind(&rack); err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), http.StatusUnprocessableEntity)
	}
	if ok, err1 := isRequestValid(&rack); !ok {
		return utils.HandleResponseIn(c, constant.FAILED, err1.Error(), http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	rack.ID = id
	rack.UpdatedAt = time.Now()
	if err := a.RackUsecase.Update(ctx, &rack); err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.CATEGORY_UPDATE_FAILED, utils.GetStatusCode(err))
	}
	return utils.HandleResponseIn(c, constant.SUCCESS, constant.CATEGORY_UPDATE_SUCCESS, http.StatusOK)
}

// Delete will delete article by given param
func (a *RackHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	err := a.RackUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
