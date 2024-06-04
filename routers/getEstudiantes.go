package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gustavosantosr/twittor/bd"
	"github.com/gustavosantosr/twittor/logger"
)

/*GetEstudiantes Leo los Apoyo */
func GetEstudiantes(w http.ResponseWriter, r *http.Request) {

	respuesta, correcto := bd.GetEstudiantes()
	if correcto != nil {
		logger.WriteLogger(fmt.Sprintf("%+v", correcto.Error()))
		http.Error(w, "Error al leer los Apoyo", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}

/*GetEstudianteByEmail Leo las Estudiante */
func GetEstudianteByEmail(w http.ResponseWriter, r *http.Request) {
	Email := r.URL.Query().Get("Email")
	respuesta, error := bd.GetEstudianteByEmail(Email)
	if error != nil {
		http.Error(w, "Error al leer los datos "+error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
