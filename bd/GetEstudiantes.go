package bd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gustavosantosr/twittor/logger"
	"github.com/gustavosantosr/twittor/models"
)

/*GetEstudiantes end point items*/
func GetEstudiantes() ([]*models.Estudiante, error) {
	err := Conexion.Ping()
	if err != nil {
		log.Printf("Error al verificar la conexión: %+v", err)
		return nil, err
	}

	query := `SELECT 
        A.EMPLID AS Emplid, 
        A.NATIONAL_ID_TYPE as TipoDocumento, 
        A.NATIONAL_ID as NumeroDocumento, 
        B.NAME_DISPLAY AS Nombre, 
        D.ADDRESS1 as Direccion, 
        E.EMAIL_ADDR as Email, 
        TRUNC((SYSDATE - C.BIRTHDATE) / 365.25) AS Edad,
        F.ACAD_PROG as CodigoPrograma, 
        G.ETY_PROG_DESCR_LG as Programa, 
        G.ETY_PROG_UNID_TOT as Ety
    FROM 
        SYSADM.PS_PERS_NID A
    LEFT JOIN 
        SYSADM.PS_ADDRESSES D ON A.EMPLID = D.EMPLID AND D.EFFDT = (SELECT MAX(D_ED.EFFDT) FROM SYSADM.PS_ADDRESSES D_ED WHERE D.EMPLID = D_ED.EMPLID AND D_ED.EFFDT <= SYSDATE)
    LEFT JOIN 
        SYSADM.PS_EMAIL_ADDRESSES E ON A.EMPLID = E.EMPLID AND E.PREF_EMAIL_FLAG = 'Y'
    LEFT JOIN 
        SYSADM.PS_PERSONAL_DATA C ON A.EMPLID = C.EMPLID
    LEFT JOIN 
        SYSADM.PS_NAMES B ON A.EMPLID = B.EMPLID AND B.NAME_TYPE = 'PRI' AND B.EFF_STATUS = 'A' AND B.EFFDT = (SELECT MAX(B_ED.EFFDT) FROM SYSADM.PS_NAMES B_ED WHERE B.EMPLID = B_ED.EMPLID AND B.NAME_TYPE = B_ED.NAME_TYPE AND B_ED.EFFDT <= SYSDATE)
    INNER JOIN 
        SYSADM.PS_ACAD_PROG F ON A.EMPLID = F.EMPLID AND F.EFFDT = (SELECT MAX(F_ED.EFFDT) FROM SYSADM.PS_ACAD_PROG F_ED WHERE F.EMPLID = F_ED.EMPLID AND F.ACAD_CAREER = F_ED.ACAD_CAREER AND F.STDNT_CAR_NBR = F_ED.STDNT_CAR_NBR AND F_ED.EFFDT <= SYSDATE) AND F.EFFSEQ = (SELECT MAX(F_ES.EFFSEQ) FROM SYSADM.PS_ACAD_PROG F_ES WHERE F.EMPLID = F_ES.EMPLID AND F.ACAD_CAREER = F_ES.ACAD_CAREER AND F.STDNT_CAR_NBR = F_ES.STDNT_CAR_NBR AND F.EFFDT = F_ES.EFFDT)
    INNER JOIN 
        SYSADM.PS_ETY_HV1_TBL G ON F.ACAD_PROG = G.ACAD_PROG AND G.INSTITUTION = F.INSTITUTION AND G.EFFDT = (SELECT MAX(G_ED.EFFDT) FROM SYSADM.PS_ETY_HV1_TBL G_ED WHERE G.INSTITUTION = G_ED.INSTITUTION AND G.ACAD_PROG = G_ED.ACAD_PROG AND G_ED.EFFDT <= SYSDATE)
    WHERE 
        A.PRIMARY_NID = 'Y'`

	log.Printf("Ejecutando consulta: %s", query)
	rows, err := Conexion.Query(query)
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	estudiantesMap := make(map[string]*models.Estudiante)

	for rows.Next() {
		var (
			emplid, tipoDocumento, numeroDocumento, nombre, direccion, email, codigoPrograma, programa string
			ety                                                                                        sql.NullInt64
			edad                                                                                       sql.NullInt64
		)
		err := rows.Scan(&emplid, &tipoDocumento, &numeroDocumento, &nombre, &direccion, &email, &edad, &codigoPrograma, &programa, &ety)
		if err != nil {
			log.Printf("Error leyendo las filas: %+v", err)
			return nil, err
		}

		estudiante, exists := estudiantesMap[emplid]
		if !exists {
			estudiante = &models.Estudiante{
				Emplid:          emplid,
				TipoDocumento:   tipoDocumento,
				NumeroDocumento: numeroDocumento,
				Nombre:          nombre,
				Direccion:       direccion,
				Email:           email,
				Edad:            edad,
				Programas:       []models.Programa{},
			}
			estudiantesMap[emplid] = estudiante
		}

		programaItem := models.Programa{
			CodigoPrograma: codigoPrograma,
			Programa:       programa,
			Ety:            ety,
		}
		estudiante.Programas = append(estudiante.Programas, programaItem)
	}

	var resultados []*models.Estudiante
	for _, estudiante := range estudiantesMap {
		resultados = append(resultados, estudiante)
	}

	return resultados, nil
}

func GetEstudianteByEmail(email string) (*models.Estudiante, error) {
	err := Conexion.Ping()
	if err != nil {
		log.Printf("Error al verificar la conexión: %+v", err)
		return nil, err
	}

	query := `SELECT 
        A.EMPLID AS Emplid, 
        A.NATIONAL_ID_TYPE as TipoDocumento, 
        A.NATIONAL_ID as NumeroDocumento, 
        B.NAME_DISPLAY AS Nombre, 
        D.ADDRESS1 as Direccion, 
        E.EMAIL_ADDR as Email, 
        TRUNC((SYSDATE - C.BIRTHDATE) / 365.25) AS Edad,
        F.ACAD_PROG as CodigoPrograma, 
        G.ETY_PROG_DESCR_LG as Programa, 
        G.ETY_PROG_UNID_TOT as Ety
    FROM 
        SYSADM.PS_PERS_NID A
    LEFT JOIN 
        SYSADM.PS_ADDRESSES D ON A.EMPLID = D.EMPLID AND D.EFFDT = (SELECT MAX(D_ED.EFFDT) FROM SYSADM.PS_ADDRESSES D_ED WHERE D.EMPLID = D_ED.EMPLID AND D_ED.EFFDT <= SYSDATE)
    LEFT JOIN 
        SYSADM.PS_EMAIL_ADDRESSES E ON A.EMPLID = E.EMPLID AND E.PREF_EMAIL_FLAG = 'Y'
    LEFT JOIN 
        SYSADM.PS_PERSONAL_DATA C ON A.EMPLID = C.EMPLID
    LEFT JOIN 
        SYSADM.PS_NAMES B ON A.EMPLID = B.EMPLID AND B.NAME_TYPE = 'PRI' AND B.EFF_STATUS = 'A' AND B.EFFDT = (SELECT MAX(B_ED.EFFDT) FROM SYSADM.PS_NAMES B_ED WHERE B.EMPLID = B_ED.EMPLID AND B.NAME_TYPE = B_ED.NAME_TYPE AND B_ED.EFFDT <= SYSDATE)
    INNER JOIN 
        SYSADM.PS_ACAD_PROG F ON A.EMPLID = F.EMPLID AND F.EFFDT = (SELECT MAX(F_ED.EFFDT) FROM SYSADM.PS_ACAD_PROG F_ED WHERE F.EMPLID = F_ED.EMPLID AND F.ACAD_CAREER = F_ED.ACAD_CAREER AND F.STDNT_CAR_NBR = F_ED.STDNT_CAR_NBR AND F_ED.EFFDT <= SYSDATE) AND F.EFFSEQ = (SELECT MAX(F_ES.EFFSEQ) FROM SYSADM.PS_ACAD_PROG F_ES WHERE F.EMPLID = F_ES.EMPLID AND F.ACAD_CAREER = F_ES.ACAD_CAREER AND F.STDNT_CAR_NBR = F_ES.STDNT_CAR_NBR AND F.EFFDT = F_ES.EFFDT)
    INNER JOIN 
        SYSADM.PS_ETY_HV1_TBL G ON F.ACAD_PROG = G.ACAD_PROG AND G.INSTITUTION = F.INSTITUTION AND G.EFFDT = (SELECT MAX(G_ED.EFFDT) FROM SYSADM.PS_ETY_HV1_TBL G_ED WHERE G.INSTITUTION = G_ED.INSTITUTION AND G.ACAD_PROG = G_ED.ACAD_PROG AND G_ED.EFFDT <= SYSDATE)
    WHERE 
        A.PRIMARY_NID = 'Y' AND E.EMAIL_ADDR = :1`

	log.Printf("Ejecutando consulta: %s", query)
	rows, err := Conexion.Query(query, email)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %+v", err)
		return nil, err
	}
	defer rows.Close()

	var estudiante *models.Estudiante

	for rows.Next() {
		var (
			emplid, tipoDocumento, numeroDocumento, nombre, direccion, emailAddr, codigoPrograma, programa string
			ety                                                                                            sql.NullInt64
			edad                                                                                           sql.NullInt64
		)
		err := rows.Scan(&emplid, &tipoDocumento, &numeroDocumento, &nombre, &direccion, &emailAddr, &edad, &codigoPrograma, &programa, &ety)
		if err != nil {
			log.Printf("Error leyendo las filas: %+v", err)
			return nil, err
		}

		if estudiante == nil {
			estudiante = &models.Estudiante{
				Emplid:          emplid,
				TipoDocumento:   tipoDocumento,
				NumeroDocumento: numeroDocumento,
				Nombre:          nombre,
				Direccion:       direccion,
				Email:           emailAddr,
				Edad:            edad, // Convertir edad de sql.NullInt64 a int
				Programas:       []models.Programa{},
			}
		}

		programaItem := models.Programa{
			CodigoPrograma: codigoPrograma,
			Programa:       programa,
			Ety:            ety, // Convertir ety de sql.NullInt64 a int
		}
		estudiante.Programas = append(estudiante.Programas, programaItem)
	}

	if estudiante == nil {
		return nil, fmt.Errorf("no se encontró ningún estudiante con el email %s", email)
	}

	return estudiante, nil
}
