package middlewares

import (
	"Praktikum/constants"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func GenerateToken(userId int) string {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenResult, _ := token.SignedString([]byte(constants.SECRET_JWT))
	return tokenResult
}

func ExtraToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenReq := c.QueryParam("token")
		_, err := jwt.Parse(tokenReq, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.SECRET_JWT), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"massage": "Unautorized",
			})
		}
		return next(c)
	}
}