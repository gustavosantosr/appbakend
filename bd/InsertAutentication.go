package bd

import (
	"fmt"
	"log"

	"github.com/gustavosantosr/twittor/logger"
)

/*InsertAutenticacion EndPoint grabar salida*/
func InsertAutenticacion() (int64, bool, error) {
	ChequeoConnection()
	logger.WriteLogger("Insertando registro")

	stmt, es := Conexion.Prepare("INSERT INTO SYSADM.PS_UCA_AUTHENTICATION_APP (DOCUMENT, CODE, REGISTRATION_DATE, STATE) VALUES (?, ?, ?, ?)")
	if es != nil {
		log.Fatal(fmt.Sprintf("Error al preparar la consulta: %+v", es))
		return 0, false, es
	}
	defer stmt.Close() // Cerrar el statement para evitar fugas de memoria

	result, er := stmt.Exec("80919446", "222222", "2025-02-05T08:19:52-05:00", "1")
	if er != nil {
		log.Fatal(fmt.Sprintf("Error al ejecutar la consulta: %+v", er))
		return 0, false, er
	}

	resultado, err := result.LastInsertId()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error al obtener LastInsertId: %+v", err))
		return 0, false, err
	}

	logger.WriteLogger(fmt.Sprintf("Registro insertado con ID: %d", resultado))
	return resultado, true, nil
}
