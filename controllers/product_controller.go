package controllers

import (
	"gozzafadillah/lib/database"
	"gozzafadillah/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllProductController(c echo.Context) error {
	product := database.GetAllProduct()
	if len(product) == 0 {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{
			"Message": "Database Empty !!",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Success Get All Product",
		"Data":    product,
	})
}

func CreateProduct(c echo.Context) error {
	temp := models.Product{}
	c.Bind(&temp)
	status := database.AddStatus(temp)
	if status != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Status not valid",
		})
	}
	if database.ExistProduct(temp.Code) == true {
		product, err := database.StoreProduct(temp)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Request not valid",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "New Product Success Added",
			"Data":    product,
		})
	}

	// cek price sama atau berubah
	err := database.CheckPriceQty(temp.Code, temp)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Request Failed",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Product Updated",
	})
}

func ProductInController(c echo.Context) error {
	proStat := models.Product_Status{}
	c.Bind(&proStat)
	err := database.ProductIn(proStat)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Request Failed",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Sukses Menambahkan Product Barang Masuk",
	})

}

func ProductOutController(c echo.Context) error {
	proStat := models.Product_Status{}
	c.Bind(&proStat)
	err := database.ProductOut(proStat)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Sukses Menambahkan Product Barang Keluar",
	})
}
