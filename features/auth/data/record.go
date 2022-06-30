package data

import (
	"fmt"
	"gozzafadillah/features/auth"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       int
	Username string
	Password string
}

// toDomain
func (u *Users) toDomain() auth.Domain {
	return auth.Domain{
		ID:       int(u.ID),
		Username: u.Username,
		Password: u.Password,
	}
}

func toDomainList(resp []Users) []auth.Domain {
	a := []auth.Domain{}
	fmt.Println("Slice core list : ", a)
	for key := range resp {
		a = append(a, resp[key].toDomain())
	}
	return a
}

func fromDomain(domain auth.Domain) *Users {
	return &Users{
		ID:       domain.ID,
		Username: domain.Username,
		Password: domain.Password,
	}
}
