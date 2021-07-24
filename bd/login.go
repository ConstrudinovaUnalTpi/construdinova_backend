package bd

import (
	"github.com/felipe1297/construdinova_backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.User, bool) {
	
	usu, encontrado, _ := ExistenciaUsuario(email)

	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu , true

}