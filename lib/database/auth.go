package database

import (
	"errors"
	"fmt"
	"gozzafadillah/config"
	"gozzafadillah/models"
)

func Register(data models.User) (models.User, error) {
	queryData := config.DB.Save(&data)
	return data, queryData.Error
}

func Login(data models.User) (models.User, error) {
	var temp int64
	var user models.User
	fmt.Println(data)
	queryData := config.DB.Model(&data).Where("username = ? AND password = ?", data.Username, data.Password).First(&user).Count(&temp)

	if temp == 0 {
		return models.User{}, errors.New("not found")
	}

	return user, queryData.Error
}
