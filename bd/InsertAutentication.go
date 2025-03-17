package bd

import (
	"database/sql"
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

	_, er := stmt.Exec("80919446", "222222", "2025-02-05T08:19:52-05:00", "1")
	if er != nil {
		log.Fatal(fmt.Sprintf("Error al ejecutar la consulta: %+v", er))
		return 0, false, er
	}
	var lastID int64 // Variable para almacenar el ID generado
	// Ejecutar la consulta y capturar el ID insertado
	err := stmt.QueryRow("80919446", "222222", "2025-02-05 08:19:52", "1", sql.Out{Dest: &lastID}).Scan(&lastID)
	if err != nil {
		log.Println(fmt.Sprintf("Error al ejecutar la consulta: %+v", err))
		return 0, false, err
	}

	logger.WriteLogger(fmt.Sprintf("Registro insertado con ID: %d", lastID))
	return lastID, true, nil
}
