package bd

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/gustavosantosr/twittor/logger"
)

/*InsertAutenticacion EndPoint grabar salida*/
func InsertAutenticacion(document string) (int64, bool, error) {
	ChequeoConnection()
	logger.WriteLogger("Insertando registro")
	code, err := generateSecureCode()
	if err != nil {
		fmt.Println("Error generando código:", err)

	}

	stmt, es := Conexion.Prepare("INSERT INTO SYSADM.PS_UCA_AUTHENTICATION_APP (DOCUMENT, CODE, REGISTRATION_DATE, STATE) VALUES (:1, :2, TO_TIMESTAMP_TZ(:3, 'YYYY-MM-DD\"T\"HH24:MI:SS TZH:TZM'), :4)")

	if es != nil {
		log.Fatal(fmt.Sprintf("Error al preparar la consulta: %+v", es))
		return 0, false, es
	}
	defer stmt.Close() // Cerrar el statement para evitar fugas de memoria
	currentTime := time.Now().Format("2006-01-02T15:04:05-07:00")
	_, er := stmt.Exec(document, code, currentTime, "1")
	if er != nil {
		log.Fatal(fmt.Sprintf("Error al ejecutar la consulta: %+v", er))
		return 0, false, er
	}

	//logger.WriteLogger(fmt.Sprintf("Registro insertado con ID: %d", lastID))
	return 0, true, nil
}

func generateSecureCode() (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(900000)) // Genera un número entre 0 y 899999
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + 100000, nil // Asegura que el número esté entre 100000 y 999999
}
