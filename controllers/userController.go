package controllers

import (
	"Praktikum/configs"
	"Praktikum/models"

	"net/http"
	"strconv"

	"github.com/labstack/echo"
)


// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User
  
	if err := configs.DB.Find(&users).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all users",
	  "users":   users,
	})
  }
  
  
  // get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	var users []models.User
  
	id, _ := strconv.Atoi(c.Param("id"))
  
	if err := configs.DB.Find(&users).Error; err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
  
	user := configs.DB.Find(&users, id)
	if user != nil {
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get users",
		"users":   user,
	  })
	}
  
	return c.JSON(http.StatusNotFound, map[string]interface{}{
	  "message": "user not found",
	})
  
}
  
  
  // create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
  
  
	if err := configs.DB.Save(&user).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new user",
	  "user":    user,
	})
}
  
  
  // delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	var users []models.User
	id, _ := strconv.Atoi(c.Param("id"))
  
	if err := configs.DB.Find(&users).Error ;err != nil{
	  return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"massage": "fail get user",
	  })
	}
	user := configs.DB.Find(&users, id)
	if user != nil{
	  configs.DB.Delete(&users, id)
	  return c.JSON(http.StatusOK, map[string]interface{}{
		"massage": "delete user success",
		"user": users,
	  })
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
	  "massage": "user not found",
	})
}
  
  
  // update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var user models.User
	var edit models.User
	c.Bind(&edit)
  
	id,_ := strconv.Atoi(c.Param("id"))
	
	if err := configs.DB.Find(&user).Error ;err != nil{
	  return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"massage": "fail get user",
	  })
	}
	configs.DB.Find(&user, id)
	user.Name = edit.Name
	user.Email = edit.Email
	user.Password = edit.Password
  
	configs.DB.Save(&user)
  
	return c.JSON(http.StatusOK, map[string]interface{}{
	  "massage": "update user success",
	  "user" : user,
	})
}