package user

import (
	"fmt"
	"go-mongodb/delivery/common"
	"go-mongodb/entities"
	userRepo "go-mongodb/repository/user"
	"net/http"

	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	repository userRepo.User
}

func New(user userRepo.User) *UserController {
	return &UserController{
		repository: user,
	}
}

func (tc UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := tc.repository.Get()

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (tc UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userRequest UserRequestFormat

		if err := c.Bind(&userRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		userS := entities.User{
			Name:     userRequest.Name,
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		response, err := tc.repository.Create(userS)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (tc UserController) GetByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Param("email")

		user, err := tc.repository.GetByEmail(email)

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (tc UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userRequest UserRequestFormat

		if err := c.Bind(&userRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		userS := entities.User{
			Name:     userRequest.Name,
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}

		oid, err := primitive.ObjectIDFromHex(userRequest.Id)
		if err != nil {
			fmt.Println(err)
		}

		err = tc.repository.Update(userS, oid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, "Success Updated Data")
	}
}

func (tc UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userRequest UserRequestFormat

		if err := c.Bind(&userRequest); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		oid, err := primitive.ObjectIDFromHex(userRequest.Id)
		if err != nil {
			fmt.Println(err)
		}

		err = tc.repository.Delete(oid)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, "Success Deleted Data")
	}
}
