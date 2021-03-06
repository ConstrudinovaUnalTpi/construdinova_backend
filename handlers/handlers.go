package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/felipe1297/construdinova_backend/middlew"
	"github.com/felipe1297/construdinova_backend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores del seteo del puerto, el handler y se pone a escuchar en el puerto*/
func Manejadores() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
