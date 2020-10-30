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

// CategoryHandler  represent the httphandler for article
type CategoryHandler struct {
	CategoryUsecase domain.CategoryUsecase
}

// NewCategoryHandler will initialize the articles/ resources endpoint
func NewCategoryHandler(e *echo.Echo, categoryUsecase domain.CategoryUsecase) {
	handler := &CategoryHandler{
		CategoryUsecase: categoryUsecase,
	}
	e.GET("/categories", handler.FetchCategory)
	e.POST("/categories", handler.Store)
	e.GET("/categories/:id", handler.GetByID)
	e.DELETE("/categories/:id", handler.Delete)
	e.PUT("/categories/:id", handler.Update)
}

// FetchCategory will fetch the categories/category based on given params
func (a *CategoryHandler) FetchCategory(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()

	listCategory, nextCursor, err := a.CategoryUsecase.Fetch(ctx, cursor, int64(num))
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.CATEGORY_GET_FAILED, utils.GetStatusCode(err))
	}

	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return utils.HandleResponseGet(c, constant.SUCCESS, constant.SUCCESS_LOAD_DATA, http.StatusOK, listCategory)
}

// GetByID will get article by given id
func (a *CategoryHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	category, err := a.CategoryUsecase.GetByID(ctx, id)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), utils.GetStatusCode(err))
	}

	return utils.HandleResponseGet(c, constant.SUCCESS_LOAD_DATA, constant.SUCCESS, http.StatusOK, category)
}

func isRequestValid(m *domain.Category) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the category by given request body
func (a *CategoryHandler) Store(c echo.Context) (err error) {
	var category domain.Category
	err = c.Bind(&category)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), http.StatusUnprocessableEntity)
	}

	var ok bool
	if ok, err = isRequestValid(&category); !ok {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), http.StatusBadRequest)
	}

	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	category.ID = utils.GenerateSHA1(category.Category, time.Now().String())

	ctx := c.Request().Context()
	err = a.CategoryUsecase.Store(ctx, &category)
	if err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), utils.GetStatusCode(err))
	}

	return c.JSON(http.StatusCreated, category)
}

// Update will update the category given with specific id
func (a *CategoryHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var category domain.Category
	if err := c.Bind(&category); err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, err.Error(), http.StatusUnprocessableEntity)
	}
	if ok, err1 := isRequestValid(&category); !ok {
		return utils.HandleResponseIn(c, constant.FAILED, err1.Error(), http.StatusBadRequest)
	}

	ctx := c.Request().Context()
	category.ID = id
	category.UpdatedAt = time.Now()
	if err := a.CategoryUsecase.Update(ctx, &category); err != nil {
		return utils.HandleResponseIn(c, constant.FAILED, constant.CATEGORY_UPDATE_FAILED, utils.GetStatusCode(err))
	}
	return utils.HandleResponseIn(c, constant.SUCCESS, constant.CATEGORY_UPDATE_SUCCESS, http.StatusOK)
}

// Delete will delete article by given param
func (a *CategoryHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	err := a.CategoryUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}