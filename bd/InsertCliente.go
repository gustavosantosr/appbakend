package bd

/*InsertAutenticacion EndPoint grabar salida*/
func InsertAutenticacion() (int64, bool, error) {
	ChequeoConnection()

	stmt, es := Conexion.Prepare("insert into SYSADM.PS_UCA_AUTHENTICATION_APP (DOCUMENT, CODE, REGISTRATION_DATE, STATE  ) values(?, ?, ?, ?, ?)")
	if es != nil {
		return 0, false, es

	}
	//fmt.Printf("grupo: %s\n", g.RazonSocial)
	result, er := stmt.Exec("80919446",
		"222222",
		"2025-02-05T08:19:52-05:00",
		"1")

	if er != nil {
		return 0, false, er
	}

	resultado, _ := result.LastInsertId()

	return resultado, true, nil
}
