package routers

import (
	"encoding/json"
	"net/http"
	"github.com/felipe1297/construdinova_backend/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Se debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro " + err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "applicartion/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}