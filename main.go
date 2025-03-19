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

	log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
	err := http.ListenAndServeTLS(":10443", certFile, keyFile, nil)

	if err != nil {
		log.Fatal("Error al iniciar servidor HTTPS:", err)
	}
}
