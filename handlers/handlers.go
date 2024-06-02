package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gustavosantosr/twittor/middlew"
	"github.com/gustavosantosr/twittor/routers"
	"github.com/rs/cors"
)

/*Manejadores manejador de urls*/
func Manejadores() {
	router := mux.NewRouter()
	/*EndPoints Terminados*/
	router.HandleFunc("/getdocentes", routers.GetDocentes).Methods("GET")
	router.HandleFunc("/getproducto", middlew.ChequeoBD(routers.GetDocentes)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
