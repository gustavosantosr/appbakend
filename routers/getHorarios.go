package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gustavosantosr/twittor/bd"
)

/*GetHorariosbyEmplid Leo las Horarios */
func GetHorariosbyEmplid(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("Emplid") == "" {
		http.Error(w, "debe enviar el Codigo", http.StatusBadRequest)
		return
	}
	Emplid := r.URL.Query().Get("Emplid")
	respuesta, error := bd.GetHorariosbyEmplid(Emplid)
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
