package auth

import "context"

type Domain struct {
	ID       int
	Username string
	Password string
}

type Usecase interface {
	CreateToken(Username, Password string) string
	Register(c context.Context) Domain
}

type Data interface {
	GetByUsername(username string) ([]Domain, error)
	InsertUser(c context.Context, data *Domain) error
}
