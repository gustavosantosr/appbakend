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

	/*Programas*/
	router.HandleFunc("/getprogramas", routers.GetProgramas).Methods("GET")
	router.HandleFunc("/getprogramasbyemplid", routers.GetProgramabyEmplid).Methods("GET")
	/*Materias*/
	router.HandleFunc("/getmaterias", routers.GetMaterias).Methods("GET")
	router.HandleFunc("/getmateriasbyemplid", routers.GetMateriabyEmplid).Methods("GET")
	/*Horarios*/
	router.HandleFunc("/gethorariosbyemplid", routers.GetHorariosbyEmplid).Methods("GET")
	/*Autenticacion*/
	router.HandleFunc("/getautenticacion", routers.GetAutenticacion).Methods("GET")
	router.HandleFunc("/getautenticacionTotal", routers.GetAutenticacionTotal).Methods("GET")
	router.HandleFunc("/getautenticacionbycode", routers.GetAutenticacionbyCode).Methods("GET")

	// Configuración más específica de CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permite cualquier origen
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	// Rutas de los certificados
	//certFile := "/etc/ssl/certs/wildcard2024.crt"
	//keyFile := "/etc/pki/tls/private/wildcard2024.key"
	handler := corsHandler.Handler(router)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", handler)

	if err != nil {
		log.Fatal("Error al iniciar servidor HTTPS:", err)
	}
	//log.Println("Servidor corriendo en el puerto", PORT)
	//log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
