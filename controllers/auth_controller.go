package controllers

import (
	"gozzafadillah/lib/database"
	"gozzafadillah/middleware"
	"gozzafadillah/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) error {
	temp := models.User{}

	c.Bind(&temp)

	user, err := database.Register(temp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Request Not Found !",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Register Success !",
		"Data":    user,
	})

}

func LoginController(c echo.Context) error {
	temp := models.User{}

	c.Bind(&temp)

	user, err := database.Login(temp)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"Message": "Login Fail!",
		})
	}
	token, err := middleware.CreateToken(user.Username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"Message": "Login Fail!",
		})
	}

	userResponse := models.UserResponse{
		Username: temp.Username,
		Token:    token,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Login Success !",
		"Data":    userResponse,
	})
}
