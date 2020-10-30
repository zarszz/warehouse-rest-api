package utils

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/zarszz/warehouse-rest-api/domain"
)

type ResponseGet struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseIn struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

func HandleResponseGet(c echo.Context, status string, message string, responseCode int, data interface{}) error {
	res := ResponseGet{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return c.JSON(responseCode, res)
}

func HandleResponseIn(c echo.Context, status string, message string, responseCode int) error {
	res := ResponseIn{
		Status:  status,
		Message: message,
	}
	return c.JSON(responseCode, res)
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
