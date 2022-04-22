package database

import (
	"errors"
	"fmt"
	"gozzafadillah/config"
	"gozzafadillah/models"
)

func GetAllProduct() []models.Product {
	var products []models.Product
	config.DB.Find(&products)
	return products
}

func ExistProduct(code string) bool {
	var product []models.Product
	config.DB.Where("code = ?", code).Find(&product)
	// tidak ketemu
	if len(product) == 0 {
		return true
	}
	// ketemu
	return false
}

func StoreProduct(data models.Product) (models.Product, error) {
	// var oldData models.Product
	// config.DB.Raw("SELECT qty FROM products WHERE code = ?", data.Code).Scan(&oldData)
	// var qty uint
	// qty = data.Qty + oldData.Qty

	newData := models.Product{
		ID:              data.ID,
		Code:            data.Code,
		Product_type_id: data.Product_type_id,
		Name:            data.Name,
		Description:     data.Description,
		Price:           data.Price,
		Qty:             data.Qty,
	}
	queryData := config.DB.Create(&newData)

	return data, queryData.Error
}

func CheckPriceQty(code string, data models.Product) error {
	var product models.Product
	config.DB.Raw("SELECT qty, price FROM products WHERE code = ?", code).Scan(&product)
	qty := data.Qty + product.Qty
	queryData := config.DB.Model(&data).Where("code = ?", code).Updates(map[string]interface{}{"price": data.Price, "qty": qty})
	return queryData.Error
}

func AddStatus(product models.Product) error {
	data := models.Product_Status{
		Product_Id:  product.ID,
		Status:      "Barang Masuk",
		Total_Price: product.Price * product.Qty,
		Qty:         product.Qty,
	}
	queryData := config.DB.Create(&data)

	return queryData.Error
}

func ProductIn(status models.Product_Status) error {
	var product models.Product
	config.DB.Where("id", status.Product_Id).Find(&product)
	fmt.Println(product.Price)
	product_stat := models.Product_Status{
		ID:          status.ID,
		Product_Id:  status.Product_Id,
		Status:      "Barang Masuk",
		Qty:         status.Qty,
		Total_Price: product.Price * status.Qty,
	}
	qty := status.Qty + product.Qty
	config.DB.Model(&product).Where("id = ? ", status.Product_Id).Update("qty", qty)
	queryData := config.DB.Create(&product_stat)

	return queryData.Error
}

func ProductOut(status models.Product_Status) error {
	var product models.Product
	config.DB.Where("id", status.Product_Id).Find(&product)
	fmt.Println(product.Qty)

	total := product.Price * status.Qty

	product_stat := models.Product_Status{
		ID:          status.ID,
		Product_Id:  status.Product_Id,
		Status:      "Barang Keluar",
		Qty:         status.Qty,
		Total_Price: (total * 20 / 100) + total,
	}
	fmt.Println("product", product.Qty, "req stat", status.Qty)
	qty := product.Qty - status.Qty
	if product.Qty < status.Qty {
		return errors.New("Inputan Anda Salah")
	}
	update := config.DB.Model(&product).Where("id = ? ", status.Product_Id).Update("qty", qty)
	if update.RowsAffected == 0 {
		return errors.New("product tidak ditemukan")
	}
	queryData := config.DB.Create(&product_stat)

	return queryData.Error
}
