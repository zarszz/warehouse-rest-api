package delivery

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/zarszz/warehouse-rest-api/constant"
	"github.com/zarszz/warehouse-rest-api/domain"
	"github.com/zarszz/warehouse-rest-api/utils"
)

type warehouseHandler struct {
	warehouseUsecase domain.WarehouseUseCase
}

func NewWarehouseHandler(e *echo.Echo, warehouseUsecase domain.WarehouseUseCase) {
	handler := &warehouseHandler{
		warehouseUsecase: warehouseUsecase,
	}
	e.POST("/warehouses", handler.Store)
	e.GET("/warehouses", handler.Fetch)
	e.GET("/warehouses/:id", handler.GetByID)
	e.GET("/warehouses/:id/rooms", handler.FetchRoom)
	e.PUT("/warehouses/:id", handler.Update)
	e.DELETE("/warehouses/:id", handler.Delete)
}

func (w *warehouseHandler) Fetch(c echo.Context) error {
	num, _ := strconv.Atoi(c.QueryParam("num"))
	ctx := c.Request().Context()

	warehouses, nextCursor, err := w.warehouseUsecase.Fetch(ctx, int64(num))
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_GET_DATA, utils.GetStatusCode(err))
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, warehouses)
}

func (w *warehouseHandler) FetchRoom(c echo.Context) error {
	warehouseID := c.Param("id")
	if warehouseID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()

	warehouses, err := w.warehouseUsecase.FetchRoom(ctx, warehouseID)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_GET_DATA, utils.GetStatusCode(err))
	}

	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, warehouses)
}

func (w *warehouseHandler) GetByID(c echo.Context) error {
	warehouseID := c.Param("id")
	if warehouseID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}
	ctx := c.Request().Context()

	warehouse, err := w.warehouseUsecase.GetByID(ctx, warehouseID)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_GET_DATA, utils.GetStatusCode(err))
	}

	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, warehouse)
}

func (w *warehouseHandler) Update(c echo.Context) error {
	warehouseID := c.Param("id")
	warehouse := new(domain.Warehouse)

	if err := c.Bind(warehouse); err != nil || warehouseID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()
	warehouse.ID = warehouseID
	warehouse.UpdatedAt = time.Now()

	err := w.warehouseUsecase.Update(ctx, warehouse)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_UPDATE_WAREHOUSE, utils.GetStatusCode(err))
	}

	return utils.HandleResponseIn(c, constant.SUCCESS, constant.SUCCESS_UPDATE_WAREHOUSE, http.StatusOK)
}

func (w *warehouseHandler) Store(c echo.Context) error {
	warehouse := new(domain.Warehouse)
	if err := c.Bind(warehouse); err != nil {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()
	warehouse.ID = utils.GenerateSHA1(warehouse.Name, time.Now().String())
	warehouse.CreatedAt = time.Now()
	warehouse.UpdatedAt = time.Now()

	id, err := w.warehouseUsecase.Store(ctx, warehouse)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_STORE_WAREHOUSE, utils.GetStatusCode(err))
	}
	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_STORE_WAREHOUSE, http.StatusOK, map[string]interface{}{
		"warehouse_id": id,
	})
}
func (w *warehouseHandler) Delete(c echo.Context) error {
	warehouseID := c.Param("id")

	if warehouseID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()

	err := w.warehouseUsecase.Delete(ctx, warehouseID)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_DELETE_WAREHOUSE, utils.GetStatusCode(err))
	}

	return utils.HandleResponseIn(c, constant.SUCCESS, constant.SUCCESS_DELETE_WAREHOUSE, http.StatusOK)
}
