package models

import "gorm.io/gorm"

type Product_type struct {
	gorm.Model
	ID   int
	Name string
}

type Product struct {
	gorm.Model
	ID              uint   `gorm:"primaryKey" json:"ID"`
	Code            string `json:"code"`
	Product_type_id uint   `json:"product_type"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Price           int    `json:"price"`
	Qty             int    `json:"qty"`
}

type Product_Status struct {
	gorm.Model
	ID          uint `gorm:"primaryKey" json:"ID"`
	Product_Id  uint `json:"product_id"`
	Status      string
	Qty         int `json:"qty"`
	Total_Price int
}
