package delivery

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/zarszz/warehouse-rest-api/constant"
	"github.com/zarszz/warehouse-rest-api/domain"
	"github.com/zarszz/warehouse-rest-api/utils"
)

type UserAddressHandler struct {
	UserAddressUsecase domain.UserAddressUsecase
}

func NewUserAddressHandler(c *echo.Echo, userAddressUsecase domain.UserAddressUsecase) {
	handler := &UserAddressHandler{
		UserAddressUsecase: userAddressUsecase,
	}
	c.POST("/users/:id/address", handler.Store)
	c.GET("/users/:id/address", handler.FetchUserAddress)
	c.PUT("/users/:id/address", handler.Update)
}

func (userAddressHandler *UserAddressHandler) Store(c echo.Context) error {
	var userAddress domain.UserAddress

	if err := c.Bind(&userAddress); err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, "Input should not empty", utils.GetStatusCode(err))
	}

	ctx := c.Request().Context()

	userAddress.UserID = c.Param("id")
	addressId := utils.GenerateSHA1(userAddress.UserID, time.Now().String(), userAddress.Address)
	userAddress.AddressID = addressId
	userAddress.CreatedAt = time.Now()
	userAddress.UpdatedAt = time.Now()

	inputError := userAddressHandler.UserAddressUsecase.Store(ctx, userAddress)
	if inputError != nil {
		return utils.HandleResponseIn(c, constant.FAILED, inputError.Error(), utils.GetStatusCode(inputError))
	}

	return utils.HandleResponseIn(c, constant.SUCCESS, "Successfully input user address data", http.StatusOK)
}

func (userAddressHandler *UserAddressHandler) FetchUserAddress(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return utils.HandleResponseIn(c, constant.FAILED, "User id should not empty", http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	userAddress, err := userAddressHandler.UserAddressUsecase.FetchUserAddress(ctx, userID)
	if err != nil {
		return utils.HandleResponseIn(c, constant.ERROR, "Error when get user address", utils.GetStatusCode(err))
	}

	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, userAddress)
}

func (userAddressHandler UserAddressHandler) Update(c echo.Context) error {
	var userAddress domain.UserAddress

	if err := c.Bind(&userAddress); err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, "Input should not empty", utils.GetStatusCode(err))
	}

	ctx := c.Request().Context()
	userAddress.UserID = c.Param("id")

	inputError := userAddressHandler.UserAddressUsecase.Update(ctx, userAddress)
	if inputError != nil {
		return utils.HandleResponseIn(c, constant.FAILED, inputError.Error(), utils.GetStatusCode(inputError))
	}

	return utils.HandleResponseIn(c, constant.SUCCESS, "Successfully input user address data", http.StatusOK)
}
