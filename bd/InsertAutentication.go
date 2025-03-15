package bd

import (
	"fmt"

	"github.com/gustavosantosr/twittor/logger"
)

/*InsertAutenticacion EndPoint grabar salida*/
func InsertAutenticacion() (int64, bool, error) {
	ChequeoConnection()
	logger.WriteLogger(fmt.Sprintf("Insertando registro"))
	stmt, es := Conexion.Prepare("insert into SYSADM.PS_UCA_AUTHENTICATION_APP (DOCUMENT, CODE, REGISTRATION_DATE, STATE  ) values(?, ?, ?, ?)")
	if es != nil {
		return 0, false, es

	}
	//fmt.Printf("grupo: %s\n", g.RazonSocial)
	result, er := stmt.Exec("80919446",
		"222222",
		"2025-02-05T08:19:52-05:00",
		"1")

	if er != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", er.Error()))
		return 0, false, er
	}

	resultado, _ := result.LastInsertId()

	return resultado, true, nil
}
