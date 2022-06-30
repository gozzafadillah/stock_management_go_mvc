package data

import (
	"context"
	"errors"
	"gozzafadillah/features/auth"

	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	Conn *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) auth.Data {
	return &mysqlAuthRepository{
		Conn: conn,
	}
}

// GetByUsername implements auth.Data
func (ar *mysqlAuthRepository) GetByUsername(username string) ([]auth.Domain, error) {
	var req []Users
	queryData := ar.Conn.Where("username = ?", username).Find(&req)

	if queryData.Error != nil {
		return []auth.Domain{}, errors.New("database broken")
	}

	return toDomainList(req), nil
}

// InsertUser implements auth.Data
func (ar *mysqlAuthRepository) InsertUser(c context.Context, data *auth.Domain) error {
	req := fromDomain(*data)

	queryData := ar.Conn.Create(req)
	if queryData.Error != nil && queryData.RowsAffected == 1 {
		return errors.New("data not found")
	}

	return nil
}
