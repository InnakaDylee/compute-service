package controllers

import (
	"Praktikum/configs"
	"Praktikum/middlewares"
	"Praktikum/models"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	var login models.Login
	c.Bind(&login)
	var users models.User
	user := configs.DB.First(&users, "name = ? AND password = ? ", login.Username, login.Password)

	if user.Error == gorm.ErrRecordNotFound{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": false,
			"message": "Email/Password Not Found",
		})
	}

	if user.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": false,
			"message": "Login Fail",
		})
	}
	token := middlewares.GenerateToken(int(users.ID))
	var response models.UserResponses
	response.Id = int(users.ID)
	response.Name = users.Name
	response.Email = users.Email
	response.Token = token
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"message": "Login Success",
		"data": response,
	})
}