package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gustavosantosr/twittor/bd"
	"github.com/gustavosantosr/twittor/logger"
)

/*GetProgramas Leo los Apoyo */
func GetProgramas(w http.ResponseWriter, r *http.Request) {

	respuesta, correcto := bd.GetProgramas()
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

/*GetProgramabyEmplid Leo las Programas */
func GetProgramabyEmplid(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("Emplid") == "" {
		http.Error(w, "debe enviar el Codigo", http.StatusBadRequest)
		return
	}
	Emplid := r.URL.Query().Get("Emplid")
	respuesta, error := bd.GetProgramasbyEmplid(Emplid)
	if error != nil {
		http.Error(w, "Error al leer los datos "+error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
