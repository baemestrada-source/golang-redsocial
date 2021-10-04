package main

import (
	"log"

	"github.com/baemestrada-source/golang-redsocial/bd"
	"github.com/baemestrada-source/golang-redsocial/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
