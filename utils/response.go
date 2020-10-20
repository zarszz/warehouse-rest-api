package utils

import (
	"github.com/labstack/echo"
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

func HandleResponseGet(c echo.Context, status string, message string, responseCode int, data interface{}) error {
	res := ResponseGet{
		Status:  status,
		Message: message,
		Data:    data,
	}
	return c.JSON(responseCode, res)
}

func HandleResponseIn(c echo.Context, status string, message string, responseCode int) error {
	res := ResponseGet{
		Status:  status,
		Message: message,
	}
	return c.JSON(responseCode, res)
}
