package bd

import (
	"fmt"

	"github.com/gustavosantosr/twittor/logger"
	"github.com/gustavosantosr/twittor/models"
)

/*GetEstudiantes end point items*/
func GetEstudiantes() ([]*models.Estudiante, error) {
	err := Conexion.Ping()
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		logger.WriteLogger(fmt.Sprintf("%+v", err.Error()))
		return nil, err
	}
	query := `SELECT 
    A.EMPLID AS Emplid, 
    A.NATIONAL_ID_TYPE as TipoDocumento, 
    A.NATIONAL_ID as NumeroDocumento, 
    B.NAME_DISPLAY AS Nombre, 
    D.ADDRESS1 as Direccion, 
    E.EMAIL_ADDR as Email, 
    F.ACAD_PROG as CodigoPrograma, 
    G.ETY_PROG_DESCR_LG as Programa, 
    G.ETY_PROG_UNID_TOT as Ety, 
    TRUNC((SYSDATE - C.BIRTHDATE) / 365.25) AS Edad
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
	logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", query))
	rows, err := Conexion.Query(query)
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Estudiante
	for rows.Next() {
		var item models.Estudiante
		err := rows.Scan(&item.Emplid,
			&item.TipoDocumento,
			&item.NumeroDocumento,
			&item.Nombre,
			&item.Direccion,
			&item.Email,
			&item.CodigoPrograma,
			&item.Programa,
			&item.Ety,
			&item.Edad)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}
