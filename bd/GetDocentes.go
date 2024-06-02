package bd

import (
	"fmt"

	"github.com/gustavosantosr/twittor/logger"
	"github.com/gustavosantosr/twittor/models"
)

/*GetDocentes end point items*/
func GetDocentes() ([]*models.Docente, error) {
	err := Conexion.Ping()
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		logger.WriteLogger(fmt.Sprintf("%+v", err.Error()))
		return nil, err
	}
	query := `SELECT DISTINCT
	C.NATIONAL_ID                            AS Cedula,
	FIRST_NAME || ' ' || MIDDLE_NAME         AS Nombres,
	LAST_NAME || ' ' || SECOND_LAST_NAME     AS Apellidos,
	(SELECT STRING_TEXT
	   FROM SYSADM.PS_STRINGS_TBL G
	  WHERE     G.PROGRAM_ID = 'ORGROUP'
			AND G.STRING_ID = F.ACAD_ORG)    AS Facultad,
	F.ACAD_ORG                               AS DeptoCod,
	F.DESCR                                  AS Departamento,
	B.INSTRUCTOR_CLASS                       AS Tipo,
	E.DESCR                                  AS Vinculacion
FROM SYSADM.PS_ETY_INST_TRM_HD  A
	LEFT JOIN SYSADM.PS_INSTRUCTOR_TERM B
		ON (    B.EMPLID = A.EMPLID
			AND B.INSTITUTION = A.INSTITUTION
			AND B.STRM = A.STRM)
	LEFT JOIN SYSADM.PS_PERS_NID C
		ON (C.EMPLID = A.EMPLID AND C.PRIMARY_NID = 'Y')
	LEFT JOIN
	(SELECT *
	   FROM SYSADM.PS_NAMES D
	  WHERE     D.EFF_STATUS = 'A'
			AND D.NAME_TYPE = 'PRI'
			AND D.EFFDT =
				(SELECT MAX (D1.EFFDT)
				   FROM SYSADM.PS_NAMES D1
				  WHERE     D1.EFF_STATUS = D.EFF_STATUS
						AND D1.NAME_TYPE = D.NAME_TYPE
						AND D1.EMPLID = D.EMPLID
						AND D1.EFFDT <= SYSDATE)) D
		ON (D.EMPLID = A.EMPLID)
	LEFT JOIN
	(SELECT *
	   FROM SYSADM.PS_INSTRUCTOR_CLAS E
	  WHERE     E.EFF_STATUS = 'A'
			AND E.EFFDT =
				(SELECT MAX (E1.EFFDT)
				   FROM SYSADM.PS_INSTRUCTOR_CLAS E1
				  WHERE     E1.EFF_STATUS = E.EFF_STATUS
						AND E1.INSTRUCTOR_CLASS = E.INSTRUCTOR_CLASS
						AND E1.INSTITUTION = E.INSTITUTION
						AND E1.EFFDT <= SYSDATE)) E
		ON (E.INSTRUCTOR_CLASS = B.INSTRUCTOR_CLASS)
	LEFT JOIN
	(SELECT F.*, G.DESCR
	   FROM SYSADM.PS_INSTR_ADVISOR F, SYSADM.PS_ACAD_ORG_TBL G
	  WHERE     F.EFFDT =
				(SELECT MAX (F1.EFFDT)
				   FROM SYSADM.PS_INSTR_ADVISOR F1
				  WHERE     F1.EFF_STATUS = F.EFF_STATUS
						AND F1.EMPLID = F.EMPLID
						AND F1.INSTITUTION = F.INSTITUTION
						AND F1.EFFDT <= SYSDATE)
			--AND G.EFF_STATUS = 'A'
			AND G.EFFDT =
				(SELECT MAX (G1.EFFDT)
				   FROM SYSADM.PS_ACAD_ORG_TBL G1
				  WHERE     G1.EFF_STATUS = G.EFF_STATUS
						AND G1.ACAD_ORG = G.ACAD_ORG
						AND G1.EFFDT <= SYSDATE)
			AND G.ACAD_ORG = F.ACAD_ORG) F
		ON (    F.EMPLID = A.EMPLID
			AND F.INSTITUTION = A.INSTITUTION
			AND F.ACAD_ORG = A.ACAD_ORG)
	INNER JOIN SYSADM.PS_INSTR_TERM_DTL H
		ON (    H.EMPLID = A.EMPLID
			AND H.INSTITUTION = A.INSTITUTION
			AND H.STRM = A.STRM)
	LEFT JOIN
	(SELECT *
	   FROM SYSADM.PS_CRSE_CATALOG I
	  WHERE     I.EFFDT =
				(SELECT MAX (I1.EFFDT)
				   FROM SYSADM.PS_CRSE_CATALOG I1
				  WHERE     I1.EFF_STATUS = I.EFF_STATUS
						AND I1.CRSE_ID = I.CRSE_ID
						AND I1.EFFDT <= SYSDATE)) I
		ON (I.CRSE_ID = H.CRSE_ID)
	LEFT JOIN SYSADM.PS_CRSE_OFFER J
		ON (    J.CRSE_ID = I.CRSE_ID
			AND J.EFFDT = I.EFFDT
			AND J.CRSE_OFFER_NBR = H.CRSE_OFFER_NBR
			AND J.ACAD_ORG = F.ACAD_ORG)
	LEFT JOIN
	(SELECT *
	   FROM SYSADM.PS_ACAD_GROUP_TBL K
	  WHERE     K.EFF_STATUS = 'A'
			AND K.EFFDT =
				(SELECT MAX (K1.EFFDT)
				   FROM SYSADM.PS_ACAD_GROUP_TBL K1
				  WHERE     K1.EFF_STATUS = K.EFF_STATUS
						AND K1.INSTITUTION = K.INSTITUTION
						AND K1.ACAD_GROUP = K.ACAD_GROUP
						AND K1.EFFDT <= SYSDATE)) K
		ON (    K.INSTITUTION = J.INSTITUTION
			AND K.ACAD_GROUP = J.ACAD_GROUP)
	LEFT JOIN SYSADM.PS_CRSE_OFFER L
		ON (    L.CRSE_ID = I.CRSE_ID
			AND L.EFFDT = I.EFFDT
			AND L.CRSE_OFFER_NBR = H.CRSE_OFFER_NBR)
	LEFT JOIN
	(SELECT *
	   FROM SYSADM.PS_ACAD_ORG_TBL M
	  WHERE     M.EFF_STATUS = 'A'
			AND M.EFFDT =
				(SELECT MAX (M1.EFFDT)
				   FROM SYSADM.PS_ACAD_ORG_TBL M1
				  WHERE     M1.ACAD_ORG = M.ACAD_ORG
						AND M1.EFF_STATUS = M.EFF_STATUS
						AND M1.EFFDT <= SYSDATE)) M
		ON (M.ACAD_ORG = L.ACAD_ORG)
	LEFT JOIN
	(SELECT *
	   FROM SYSADM.PS_ACAD_GROUP_TBL N
	  WHERE     N.EFF_STATUS = 'A'
			AND N.EFFDT =
				(SELECT MAX (N1.EFFDT)
				   FROM SYSADM.PS_ACAD_GROUP_TBL N1
				  WHERE     N1.EFF_STATUS = N.EFF_STATUS
						AND N1.INSTITUTION = N.INSTITUTION
						AND N1.ACAD_GROUP = N.ACAD_GROUP
						AND N1.EFFDT <= SYSDATE)) N
		ON (N.ACAD_GROUP = L.ACAD_GROUP)
	LEFT JOIN SYSADM.PS_CLASS_TBL Q
		ON (    Q.CRSE_ID = H.CRSE_ID
			AND Q.CRSE_OFFER_NBR = H.CRSE_OFFER_NBR
			AND Q.STRM = H.STRM
			AND Q.SESSION_CODE = H.SESSION_CODE
			AND Q.CLASS_NBR = H.CLASS_NBR)
	LEFT JOIN SYSADM.PS_TERM_VAL_TBL R ON (R.STRM = Q.STRM)
WHERE A.STRM IN ('1820')`
	logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", query))
	rows, err := Conexion.Query(query)
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Docente
	for rows.Next() {
		var item models.Docente
		err := rows.Scan(&item.Cedula,
			&item.Nombres,
			&item.Apellidos,
			&item.Facultad,
			&item.DeptoCod,
			&item.Departamento,
			&item.Tipo,
			&item.Vinculacion)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}
