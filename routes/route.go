package routes

import (
	"Praktikum/controllers"
	"Praktikum/middlewares"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)


func Init() {
  // create a new echo instance
  e := echo.New()
  // Route / to handler function
  middlewares.LogMiddleware(e)

  e.POST("/users/login",controllers.Login)
  e.POST("/users", controllers.CreateUserController)

  g := e.Group("")
  g.Use(middlewares.ExtraToken)
  g.GET("/users", controllers.GetUsersController)
  g.GET("/users/:id", controllers.GetUserController)
  g.DELETE("/users/:id", controllers.DeleteUserController)
  g.PUT("/users/:id", controllers.UpdateUserController)
  g.GET("/books", controllers.GetBooksController)
  g.GET("/books/:id", controllers.GetBookController)
  g.POST("/books", controllers.CreateBookController)
  g.DELETE("/books/:id", controllers.DeleteBookController)
  g.PUT("/books/:id", controllers.UpdateBookController)

  // start the server, and log if it fails
  e.Logger.Fatal(e.Start(":8080"))
}