package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/felipe1297/construdinova_backend/models"
	"github.com/felipe1297/construdinova_backend/bd"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error){
	
	privateKey := []byte("construdinova-ApiRest-Backend-PrivateKey")

	claims := &models.Claim{}

	splitToken := strings.Split(tk,"Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato invalido del token")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
		return privateKey, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ExistenciaUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}