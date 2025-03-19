package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	_ "github.com/godror/godror"
	"github.com/gustavosantosr/twittor/bd"
	"github.com/gustavosantosr/twittor/handlers"
)

func main() {
	// Conectar a la base de datos
	bd.ConectarBD()

	// Información del sistema
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCHITECTURE\t", runtime.GOARCH)
	fmt.Println("CPUS\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// Iniciar manejadores de la API
	handlers.Manejadores()

	// Rutas de los certificados
	certFile := "/etc/ssl/certs/wildcard2024.crt"
	keyFile := "/etc/pki/tls/private/wildcard2024.key"

	// Iniciar el servidor HTTPS
	log.Println("Servidor HTTPS corriendo en https://0.0.0.0:8080")
	err := http.ListenAndServeTLS(":8080", certFile, keyFile, nil)
	if err != nil {
		log.Fatal("Error al iniciar servidor HTTPS:", err)
	}
}
