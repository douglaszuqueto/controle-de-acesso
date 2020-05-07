package services

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	model "github.com/douglaszuqueto/controle-de-acesso/app/models"
)

const jwtSecret = "dz"

// Auth auth
func Auth(email, password string) (*model.Admin, error) {
	res, err := model.FindAdmin(email, password)

	if err != nil {
		return nil, err
	}

	res.Token = createTokenString(res)

	return res, nil
}

func createTokenString(admin *model.Admin) string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &model.Admin{
		Name:  admin.Name,
		Email: admin.Email,
		ID:    admin.ID})

	tokenstring, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}
