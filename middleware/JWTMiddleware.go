package middleware

import (
	"gozzafadillah/constant"

	"github.com/golang-jwt/jwt"
)

func CreateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["username"] = username
	/*
		Bila ingin membatas kadaluarsa
		 claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	*/

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_JWT))
}
