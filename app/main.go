package main

import (
	"go-mongodb/config"

	_userRepo "go-mongodb/repository/user"

	_userController "go-mongodb/delivery/controllers/user"

	"go-mongodb/util"

	"go-mongodb/delivery/router"

	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MongoDBConnect(config)

	//initiate user model
	userRepo := _userRepo.New(db)

	//initiate user controller
	userController := _userController.New(userRepo)

	//create echo http
	e := echo.New()

	//register API path and controller
	router.RegisterPath(e, userController)

	// run server
	address := fmt.Sprintf(":%d", config.Port)

	if err := e.Start(address); err != nil {
		log.Info("shutting down the server")
	}
}
