package routers

import (
	"encoding/json"
	"net/http"
	"github.com/felipe1297/construdinova_backend/bd"
	"github.com/felipe1297/construdinova_backend/models"
)


func Registro(w http.ResponseWriter, r *http.Request){

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en la data recibida " + err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contrseña debe tener al menos 6 caracteres. ", 400)
		return
	}

	/*Se confirma si el usuarion no existe*/
	_,encontrado,_ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Ya existe el usuario con el email:  " + t.Email + " . ", 400)
		return
	}

	_,status, err := bd.InsertarRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un errpr al intentar realizar el registro del usuario " + err.Error() + ". ", 400)
		return
	}

	if !status {
		http.Error(w, "No se almaceno el usuario correctamente. ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}