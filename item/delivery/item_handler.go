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

type itemHandler struct {
	itemUsecase domain.ItemUseCase
}

func NewItemHandler(e *echo.Echo, itemUsecase domain.ItemUseCase) {
	handler := &itemHandler{
		itemUsecase: itemUsecase,
	}
	e.POST("/items", handler.Store)
	e.GET("/items", handler.Fetch)
	e.GET("/items/:id", handler.GetByID)
	e.PUT("/items/:id", handler.Update)
	e.DELETE("/items/:id", handler.Delete)
}

func (r *itemHandler) Fetch(c echo.Context) error {
	num, _ := strconv.Atoi(c.QueryParam("num"))
	ctx := c.Request().Context()

	items, nextCursor, err := r.itemUsecase.Fetch(ctx, int64(num))
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_GET_DATA, utils.GetStatusCode(err))
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, items)
}

func (r *itemHandler) GetByID(c echo.Context) error {
	roomID := c.Param("id")
	if roomID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}
	ctx := c.Request().Context()

	item, err := r.itemUsecase.GetByID(ctx, roomID)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_GET_DATA, utils.GetStatusCode(err))
	}

	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, item)
}

func (r *itemHandler) Update(c echo.Context) error {
	itemID := c.Param("id")
	item := new(domain.Item)

	if err := c.Bind(item); err != nil || itemID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()
	item.ID = itemID
	item.UpdatedAt = time.Now()

	err := r.itemUsecase.Update(ctx, item)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_UPDATE_ROOM, utils.GetStatusCode(err))
	}

	return utils.HandleResponseIn(c, constant.SUCCESS, constant.SUCCESS_UPDATE_ROOM, http.StatusOK)
}

func (r *itemHandler) Store(c echo.Context) error {
	item := new(domain.Item)
	if err := c.Bind(item); err != nil {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()
	item.ID = utils.GenerateSHA1(item.ItemName, item.Description, time.Now().String())
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	err := r.itemUsecase.Store(ctx, item)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_STORE_ROOM, utils.GetStatusCode(err))
	}
	return utils.HandleResponseIn(c, constant.SUCCESS, constant.SUCCESS_STORE_ROOM, http.StatusOK)
}
func (r *itemHandler) Delete(c echo.Context) error {
	itemID := c.Param("id")

	if itemID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()

	err := r.itemUsecase.Delete(ctx, itemID)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_DELETE_ROOM, utils.GetStatusCode(err))
	}

	return utils.HandleResponseIn(c, constant.SUCCESS, constant.SUCCESS_DELETE_ROOM, http.StatusOK)
}
