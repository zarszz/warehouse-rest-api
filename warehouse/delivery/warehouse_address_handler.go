package delivery

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/zarszz/warehouse-rest-api/constant"
	"github.com/zarszz/warehouse-rest-api/domain"
	"github.com/zarszz/warehouse-rest-api/utils"
)

type WarehouseAddressHandler struct {
	WarehouseAddressUsecase domain.WarehouseAddressUsecase
}

func NewUserAddressHandler(c *echo.Echo, warehouseAddressUsecase domain.WarehouseAddressUsecase) {
	handler := &WarehouseAddressHandler{
		WarehouseAddressUsecase: warehouseAddressUsecase,
	}
	c.POST("/warehouses/:id/address", handler.Store)
	c.GET("/warehouses/:id/address", handler.FetchWarehouseAddress)
	c.PUT("/warehouses/:id/address", handler.Update)
}

func (w *WarehouseAddressHandler) Store(c echo.Context) error {
	var warehouseAddress domain.WarehouseAddress

	if err := c.Bind(&warehouseAddress); err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, "Input should not empty", utils.GetStatusCode(err))
	}

	ctx := c.Request().Context()

	warehouseAddress.WarehouseID = c.Param("id")
	addressId := utils.GenerateSHA1(warehouseAddress.Address, time.Now().String(), warehouseAddress.PostalCode)
	warehouseAddress.AddressID = addressId
	warehouseAddress.CreatedAt = time.Now()
	warehouseAddress.UpdatedAt = time.Now()

	inputError := w.WarehouseAddressUsecase.Store(ctx, warehouseAddress)
	if inputError != nil {
		return utils.HandleResponseIn(c, constant.FAILED, inputError.Error(), utils.GetStatusCode(inputError))
	}

	return utils.HandleResponseIn(c, constant.SUCCESS, "Successfully input warehouse address data", http.StatusOK)
}

func (w *WarehouseAddressHandler) FetchWarehouseAddress(c echo.Context) error {
	warehouseID := c.Param("id")
	if warehouseID == "" {
		return utils.HandleResponseIn(c, constant.FAILED, "Warehouse id should not empty", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	warehouseAddress, err := w.WarehouseAddressUsecase.FetchWarehouseAddress(ctx, warehouseID)
	if err != nil {
		return utils.HandleResponseIn(c, constant.ERROR, "Error when get warehouse address", utils.GetStatusCode(err))
	}

	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, warehouseAddress)
}

func (w *WarehouseAddressHandler) Update(c echo.Context) error {
	var warehouseAddress domain.WarehouseAddress

	if err := c.Bind(&warehouseAddress); err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, "Input should not empty", utils.GetStatusCode(err))
	}

	ctx := c.Request().Context()
	warehouseAddress.WarehouseID = c.Param("id")
	warehouseAddress.UpdatedAt = time.Now()

	inputError := w.WarehouseAddressUsecase.Update(ctx, warehouseAddress)
	if inputError != nil {
		return utils.HandleResponseIn(c, constant.FAILED, inputError.Error(), utils.GetStatusCode(inputError))
	}

	return utils.HandleResponseIn(c, constant.SUCCESS, "Successfully input warehouse address data", http.StatusOK)
}
