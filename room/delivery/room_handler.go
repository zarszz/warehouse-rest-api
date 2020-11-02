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

type roomHandler struct {
	roomUsecase domain.RoomUseCase
}

func NewRoomHandler(e *echo.Echo, roomUsecase domain.RoomUseCase) {
	handler := &roomHandler{
		roomUsecase: roomUsecase,
	}
	e.POST("/rooms", handler.Store)
	e.GET("/rooms", handler.Fetch)
	e.GET("/rooms/:id", handler.GetByID)
	e.PUT("/rooms/:id", handler.Update)
	e.DELETE("/rooms/:id", handler.Delete)
}

func (r *roomHandler) Fetch(c echo.Context) error {
	num, _ := strconv.Atoi(c.QueryParam("num"))
	ctx := c.Request().Context()

	rooms, nextCursor, err := r.roomUsecase.Fetch(ctx, int64(num))
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_GET_DATA, utils.GetStatusCode(err))
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, rooms)
}

func (r *roomHandler) GetByID(c echo.Context) error {
	roomID := c.Param("id")
	if roomID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}
	ctx := c.Request().Context()

	room, err := r.roomUsecase.GetByID(ctx, roomID)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_GET_DATA, utils.GetStatusCode(err))
	}

	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, room)
}

func (r *roomHandler) Update(c echo.Context) error {
	roomID := c.Param("id")
	room := new(domain.Room)

	if err := c.Bind(room); err != nil || roomID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()
	room.ID = roomID
	room.UpdatedAt = time.Now()

	err := r.roomUsecase.Update(ctx, room)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_UPDATE_ROOM, utils.GetStatusCode(err))
	}

	return utils.HandleResponseIn(c, constant.SUCCESS, constant.SUCCESS_UPDATE_ROOM, http.StatusOK)
}

func (r *roomHandler) Store(c echo.Context) error {
	room := new(domain.Room)
	if err := c.Bind(room); err != nil {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()
	room.ID = utils.GenerateSHA1(room.Name, time.Now().String())
	room.CreatedAt = time.Now()
	room.UpdatedAt = time.Now()

	err := r.roomUsecase.Store(ctx, room)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_STORE_ROOM, utils.GetStatusCode(err))
	}
	return utils.HandleResponseIn(c, constant.SUCCESS, constant.SUCCESS_STORE_ROOM, http.StatusOK)
}
func (r *roomHandler) Delete(c echo.Context) error {
	roomID := c.Param("id")

	if roomID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.FAILED_GET_DATA, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	ctx := c.Request().Context()

	err := r.roomUsecase.Delete(ctx, roomID)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.FAILED_DELETE_ROOM, utils.GetStatusCode(err))
	}

	return utils.HandleResponseIn(c, constant.SUCCESS, constant.SUCCESS_DELETE_ROOM, http.StatusOK)
}
