package main

import (
	"fmt"
	
	"runtime"
	"sync"
	_ "github.com/godror/godror"
	"github.com/gustavosantosr/twittor/bd"
)

var wg, wg2, wg3, wg4 sync.WaitGroup

func main() {

	
	bd.ConectarBD()
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCHITECTURE\t", runtime.GOARCH)
	fmt.Println("CPUS\t", runtime.NumCPU())
	fmt.Println("Gorutines\t", runtime.NumGoroutine())

	pruebas, err := bd.GetConsolidados()
	if err != nil {
		fmt.Println("Error al obtener consolidados:", err)
		return
	}

	fmt.Printf("%+v\n", pruebas)
	
}
