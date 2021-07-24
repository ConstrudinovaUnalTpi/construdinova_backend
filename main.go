package main

import (
	"log"
	"github.com/felipe1297/construdinova_backend/bd"
	"github.com/felipe1297/construdinova_backend/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
