package router

import (
	"go-mongodb/delivery/controllers/user"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	userController *user.UserController) {

	// user
	e.GET("/user", userController.Get())
	e.GET("/user/:email", userController.GetByEmail())
	e.POST("/user", userController.Create())
	e.PUT("/user", userController.Update())
	e.DELETE("/user", userController.Delete())
}
