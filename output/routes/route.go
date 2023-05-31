package routes

import (
	"crud-with-go/controllers"

	"github.com/labstack/echo/v4"
)



func UserRoute(e *echo.Echo){
	// All relavant routes
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:userId", controllers.GetUser)
	e.PUT("/user/:userId", controllers.EditUser)
	e.DELETE("/user/:userId", controllers.DeleteUser)
}