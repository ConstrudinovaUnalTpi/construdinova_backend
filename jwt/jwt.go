package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/felipe1297/construdinova_backend/models"
)

func GenerarJWT(t models.User) (string, error){

	privateKey := []byte("construdinova-ApiRest-Backend-PrivateKey")

	payload := jwt.MapClaims{
		"email": t.Email,
		"typeUser": t.TypeUser,
		"_id": t.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr , err := token.SignedString(privateKey)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr,nil
}