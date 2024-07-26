package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gustavosantosr/twittor/routers"
	"github.com/rs/cors"
)

/*Manejadores manejador de urls*/
func Manejadores() {
	router := mux.NewRouter()
	/*EndPoints Docentes*/
	router.HandleFunc("/getdocentes", routers.GetDocentes).Methods("GET")
	/*EndPoints Estudiantes*/
	router.HandleFunc("/getestudiantes", routers.GetEstudiantes).Methods("GET")
	router.HandleFunc("/getestudiantebyemail", routers.GetEstudianteByEmail).Methods("GET")
	router.HandleFunc("/getmaterias", routers.GetMaterias).Methods("GET")
	/*Programas*/
	router.HandleFunc("/getprogramas", routers.GetProgramas).Methods("GET")
	router.HandleFunc("/getprogramasbyemplid", routers.GetProgramabyEmplid).Methods("GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
