package bd

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/gustavosantosr/twittor/logger"
)

/*Conexion es el objeto de conexión a la BD */
var (
	once     sync.Once
	Conexion *sql.DB
)

/*ConectarBD es la función que me permite conectar la BD */
func ConectarBD() *sql.DB {
	var err error
	once.Do(func() {
		//172.40.2.226
		//Conexion, err = sql.Open("godror", "gasantosr/wr3eBjbECe@//172.40.2.226:1521/PRODELEC?&standaloneConnection=1&connectionClass=POOLED&max=6000")
		//Conexion, err = sql.Open("godror", "GS80819446/GS80819446@132.145.143.45:1521/CS92PROD.subnetpsprod03.ucaldas.oraclevcn.com")
		Conexion, err = sql.Open("godror", "GS80819446/GS80819446@132.145.143.45:1521/CS92LOAD.subnetpsprod03.ucaldas.oraclevcn.com")

		if err != nil {
			logger.WriteLogger(fmt.Sprintf("%+v", err.Error()))
			log.Fatal("Open connection failed:", err.Error())
		}
		Conexion.SetMaxIdleConns(0)
		Conexion.SetMaxOpenConns(6000)

		//Conexion.SetConnMaxLifetime(0)
		//Conexion.SetMaxOpenConns(3000)

		//Conexion.SetConnMaxLifetime(0)

		fmt.Printf("Connectedo\n")
	})

	// Connect to database prod
	logger.WriteLogger(fmt.Sprintf("%+v", Conexion.Stats().OpenConnections))

	// defer conn.Close()
	return Conexion
}

/*ChequeoConnection es el Ping a la BD */
func ChequeoConnection() int {
	err := Conexion.Ping()
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("%+v", err.Error()))
		return 0
	}
	return 1
}
