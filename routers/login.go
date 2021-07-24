package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/felipe1297/construdinova_backend/bd"
	"github.com/felipe1297/construdinova_backend/jwt"
	"github.com/felipe1297/construdinova_backend/models"
)

/*Funcion para realizar el login*/
func Login(w http.ResponseWriter, r * http.Request){
	w.Header().Add("content-type","application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos " + err.Error(),400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es necesario",400)
		return
	}

	documento, existe := bd.Login(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o contraseña invalidos.",400)
		return
	}

	jwtKey, errJwt := jwt.GenerarJWT(documento)

	if errJwt != nil {
		http.Error(w, "Ocurrio un error al generar el token.",400)
		return
	}

	resp := models.TokenResponse {
		Token : jwtKey,
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:"token",
		Value: jwtKey,
		Expires: expirationTime,
	})

}