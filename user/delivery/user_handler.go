package delivery

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/zarszz/warehouse-rest-api/constant"
	"github.com/zarszz/warehouse-rest-api/domain"
	"github.com/zarszz/warehouse-rest-api/utils"
	"gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

// NewUserHandler will initialize the articles/ resources endpoint
func NewUserHandler(e *echo.Echo, userUsecase domain.UserUseCase) {
	handler := &UserHandler{
		UserUseCase: userUsecase,
	}
	e.GET("/users", handler.FetchUser)
	e.POST("/users", handler.Store)
	e.GET("/users/:id", handler.GetByID)
	e.DELETE("/users/:id", handler.Delete)
	e.PUT("/users/:id", handler.Update)
}

// FetchUser will fetch the categories/category based on given params
func (a *UserHandler) FetchUser(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	ctx := c.Request().Context()

	users, nextCursor, err := a.UserUseCase.Fetch(ctx, int64(num))
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), utils.GetStatusCode(err))
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, users)
}

// GetByID will get user by given id
func (a *UserHandler) GetByID(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return utils.HandleResponseGet(c, constant.FAILED, constant.USER_GET_FAILED, http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}
	ctx := c.Request().Context()

	user, err := a.UserUseCase.GetByID(ctx, userID)
	if err != nil {
		return utils.HandleResponseGet(c, constant.FAILED, constant.CATEGORY_GET_FAILED, utils.GetStatusCode(err), err.Error())
	}

	return utils.HandleResponseGet(c, constant.SUCCESS_LOAD_DATA, constant.SUCCESS, http.StatusOK, user)
}

func isRequestValid(m *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the user by given request body
func (a *UserHandler) Store(c echo.Context) (err error) {
	var user domain.User

	err = c.Bind(&user)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), http.StatusUnprocessableEntity)
	}
	var ok bool
	if ok, err = isRequestValid(&user); !ok {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), http.StatusBadRequest)
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	id := utils.GenerateSHA1(user.FirstName, user.LastName, time.Now().String())
	user.ID = id

	ctx := c.Request().Context()
	err = a.UserUseCase.Store(ctx, &user)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), utils.GetStatusCode(err))
	}

	return c.JSON(http.StatusCreated, user)
}

// Update will update the user given with specific id
func (a *UserHandler) Update(c echo.Context) error {
	userID := c.Param("id")

	if userID == "" {
		return utils.HandleResponseIn(c, constant.FAILED, domain.ErrBadParamInput.Error(), http.StatusUnprocessableEntity)
	}

	var user domain.User
	if err := c.Bind(&user); err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), http.StatusUnprocessableEntity)
	}

	user.ID = userID

	if ok, err1 := isRequestValid(&user); !ok {
		return utils.HandleResponseIn(c, constant.FAILED, err1.Error(), http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	if err := a.UserUseCase.Update(ctx, &user); err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.USER_UPDATE_FAILED, utils.GetStatusCode(err))
	}
	return utils.HandleResponseIn(c, constant.SUCCESS, constant.USER_UPDATE_SUCCESS, http.StatusOK)
}

// Delete will delete user by given param
func (a *UserHandler) Delete(c echo.Context) error {
	userID := c.Param("id")
	if userID == "" {
		return utils.HandleResponseIn(c, constant.USER_GET_FAILED, domain.ErrBadParamInput.Error(), http.StatusBadRequest)
	}

	ctx := c.Request().Context()

	err := a.UserUseCase.Delete(ctx, userID)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
