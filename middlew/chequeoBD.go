package middlew

import (
	"net/http"

	"github.com/baemestrada-source/golang-redsocial/bd"
)

/*ChequeoBD es el midlware que permite conocer el estado de la bd*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
