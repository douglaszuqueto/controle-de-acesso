package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

// Admin struct
type Admin struct {
	ID       string
	Name     string
	Email    string
	Gravatar string
	Token    string
	jwt.StandardClaims
}

// Admins admins
type Admins []Admin

// FindAdmin findAdmin
func FindAdmin(email, password string) (*Admin, error) {

	if validate.Var(email, "required,email") != nil {
		return nil, errors.New("Email obrigatório")
	}

	if validate.Var(password, "required") != nil {
		return nil, errors.New("Senha obrigatória")
	}

	stmt, err := db.Prepare("SELECT id, name, email FROM sistema.admin WHERE email=$1 AND password=$2")
	defer stmt.Close()

	checkErr(err)

	admin := Admin{}
	err = stmt.QueryRow(email, password).Scan(&admin.ID, &admin.Name, &admin.Email)

	if err != nil {
		return nil, errors.New("Usuário inválido")
	}

	hash := md5.New()
	hash.Write([]byte(admin.Email))

	admin.Gravatar = fmt.Sprintf("https://secure.gravatar.com/avatar/%s", hex.EncodeToString(hash.Sum(nil)))

	return &admin, nil
}
