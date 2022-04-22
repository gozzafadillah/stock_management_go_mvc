package routes

import (
	"gozzafadillah/constant"
	"gozzafadillah/controllers"
	m "gozzafadillah/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// Log
	m.LogMiddleware(e)

	// auth
	routeAuth := e.Group("")
	routeAuth.POST("/register", controllers.RegisterController)
	routeAuth.POST("/login", controllers.LoginController)

	// product
	productJWT := e.Group("")
	// JWT init
	productJWT.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

	productJWT.GET("/product", controllers.GetAllProductController)
	productJWT.POST("/product", controllers.CreateProduct)
	productJWT.POST("/product/in", controllers.ProductInController)
	productJWT.POST("/product/out", controllers.ProductOutController)
	return e
}
