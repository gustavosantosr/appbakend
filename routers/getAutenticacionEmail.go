package routers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gustavosantosr/twittor/bd"
	"github.com/gustavosantosr/twittor/logger"
)

/*GetAutenticacion Leo los Apoyo */
func GetAutenticacionbyCode(w http.ResponseWriter, r *http.Request) {
	Code := r.URL.Query().Get("Code")
	Document := r.URL.Query().Get("Document")

	//1053846544
	respuesta, correcto := bd.GetEmailByCredentials(Code, Document)
	//_, respuesta, correcto := bd.InsertAutenticacion()
	if correcto != nil {
		logger.WriteLogger(fmt.Sprintf("%+v", correcto.Error()))
		log.Fatal(fmt.Sprintf("%+v", correcto.Error()))
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
