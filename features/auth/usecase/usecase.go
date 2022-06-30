package usecase

import (
	"context"
	"gozzafadillah/features/auth"

	"github.com/go-playground/validator/v10"
)

type userUsecase struct {
	userData auth.Data
	validate *validator.Validate
}

func NewAuthUsecase(ud auth.Data) auth.Usecase {
	return &userUsecase{
		userData: ud,
		validate: validator.New(),
	}
}

// CreateToken implements auth.Usecase
func (*userUsecase) CreateToken(Username string, Password string) string {
	panic("unimplemented")
}

// Register implements auth.Usecase
func (*userUsecase) Register(c context.Context) auth.Domain {
	panic("unimplemented")
}
