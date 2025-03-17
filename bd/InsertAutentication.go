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

	stmt, es := Conexion.Prepare("INSERT INTO SYSADM.PS_UCA_AUTHENTICATION_APP (DOCUMENT, CODE, REGISTRATION_DATE, STATE) VALUES (:1, :2, TO_TIMESTAMP_TZ(:3, 'YYYY-MM-DD\"T\"HH24:MI:SS TZH:TZM'), :4)")

	if es != nil {
		log.Fatal(fmt.Sprintf("Error al preparar la consulta: %+v", es))
		return 0, false, es
	}
	defer stmt.Close() // Cerrar el statement para evitar fugas de memoria

	_, er := stmt.Exec("80919442", "2222222", "2025-02-05T08:19:52-05:00", "2")
	if er != nil {
		log.Fatal(fmt.Sprintf("Error al ejecutar la consulta: %+v", er))
		return 0, false, er
	}

	//logger.WriteLogger(fmt.Sprintf("Registro insertado con ID: %d", lastID))
	return 0, true, nil
}
