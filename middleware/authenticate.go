package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/zarszz/warehouse-rest-api/auth"
	"github.com/zarszz/warehouse-rest-api/constant"
	"github.com/zarszz/warehouse-rest-api/utils"
)

func JWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func (c echo.Context) error {
			header := c.Request().Header
			
			if !strings.Contains(header.Get("Authorization"), "Bearer") {
				return utils.HandleResponseIn(c, constant.FAILED, "Not token provided", http.StatusBadRequest)
			}
			
			tokenString := strings.Split(header.Get("Authorization"), " ")[1]
			token, err := auth.NewService().ValidateToken(tokenString)
			if err != nil {
				return utils.HandleResponseIn(c, constant.FAILED, "Token not authorizaed", http.StatusUnauthorized)
			}
				
			claim, ok := token.Claims.(jwt.MapClaims); 
			if !ok && !token.Valid {
				return utils.HandleResponseIn(c, constant.FAILED, "Token not authorizaed", http.StatusUnauthorized)
			}
	
			userId := claim["user_id"]
			c.Set("user", userId)
	
			return next(c)
		}
	}
}
