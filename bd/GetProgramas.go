package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/twittor/logger"
	"github.com/gustavosantosr/twittor/models"
)

/*GetProgramas end point items*/
func GetProgramas() ([]*models.Programas, error) {
	err := Conexion.Ping()
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		logger.WriteLogger(fmt.Sprintf("%+v", err.Error()))
		return nil, err
	}
	query := `SELECT 
    DISTINCT A.ACAD_CAREER AS TipoCarrera,
    A.ACAD_PROG AS CodAcademicoPrograma,
    TO_CHAR(C.ETY_PROG_DESCR_LG) AS ProgramaDescripcion, 
    TO_CHAR(C.ETY_PROG_DURACION) AS ProgramaDuracion,
    TO_CHAR(C.ETY_PROG_PERF_APL) AS ProgramaPerfilAplicacion,
    TO_CHAR(C.ETY_PROG_PERF_PROF) AS ProgramaPerfilProfesional,
    TO_CHAR(C.ETY_PROG_PERF_PROF) AS ProgramaCompromisoProfesional,
    TO_CHAR(C.ETY_PROG_PERF_PROF) AS ProgramaPerfilOcupacional
FROM 
    SYSADM.PS_ACAD_PROG A
    JOIN SYSADM.PS_ETY_HV1_TBL C 
        ON C.INSTITUTION = A.INSTITUTION
        AND C.ACAD_PROG = A.ACAD_PROG
        AND A.PROG_STATUS = 'AC'`
	logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", query))
	rows, err := Conexion.Query(query)
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Programas
	for rows.Next() {
		var item models.Programas
		err := rows.Scan(&item.TipoCarrera,
			&item.CodAcademicoPrograma,
			&item.ProgramaDescripcion,
			&item.ProgramaDuracion,
			&item.ProgramaPerfilAplicacion,
			&item.ProgramaPerfilProfesional,
			&item.ProgramaCompromisoProfesional,
			&item.ProgramaPerfilOcupacional)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

/*GetProgramasbyEmplid end point items*/
func GetProgramasbyEmplid(t string) ([]*models.ProgramasbyEmplid, error) {
	err := Conexion.Ping()
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		logger.WriteLogger(fmt.Sprintf("%+v", err.Error()))
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	fmt.Println("parametro ts\t", t)
	query := `
  SELECT TO_CHAR(A.EMPLID) as Emplid, 
TO_CHAR(A.ACAD_CAREER) as TipoCarrera, 
TO_CHAR(A.PROG_STATUS) as Status, 
TO_CHAR(B.ACAD_PLAN) as CodPlanAcademico, 
TO_CHAR(C.ETY_PROG_DESCR_LG) as ProgramaDescripcion,
TO_CHAR(C.ETY_PROG_DURACION) AS ProgramaDuracion 
  FROM SYSADM.PS_ACAD_PROG A
  JOIN SYSADM.PS_ACAD_PLAN B ON A.EMPLID = B.EMPLID
  JOIN SYSADM.PS_ETY_HV1_TBL C ON C.INSTITUTION = A.INSTITUTION AND C.ACAD_PROG = A.ACAD_PROG
  WHERE A.EFFDT = (SELECT MAX(A_ED.EFFDT) FROM SYSADM.PS_ACAD_PROG A_ED
                   WHERE A.EMPLID = A_ED.EMPLID
                     AND A.ACAD_CAREER = A_ED.ACAD_CAREER
                     AND A.STDNT_CAR_NBR = A_ED.STDNT_CAR_NBR
                     AND A_ED.EFFDT <= SYSDATE)
  AND A.EFFSEQ = (SELECT MAX(A_ES.EFFSEQ) FROM SYSADM.PS_ACAD_PROG A_ES
                  WHERE A.EMPLID = A_ES.EMPLID
                    AND A.ACAD_CAREER = A_ES.ACAD_CAREER
                    AND A.STDNT_CAR_NBR = A_ES.STDNT_CAR_NBR
                    AND A.EFFDT = A_ES.EFFDT)
  AND B.EFFDT = (SELECT MAX(B_ED.EFFDT) FROM SYSADM.PS_ACAD_PLAN B_ED
                 WHERE B.EMPLID = B_ED.EMPLID
                   AND B.ACAD_CAREER = B_ED.ACAD_CAREER
                   AND B.STDNT_CAR_NBR = B_ED.STDNT_CAR_NBR
                   AND B_ED.EFFDT <= SYSDATE)
  AND B.EFFSEQ = (SELECT MAX(B_ES.EFFSEQ) FROM SYSADM.PS_ACAD_PLAN B_ES
                  WHERE B.EMPLID = B_ES.EMPLID
                    AND B.ACAD_CAREER = B_ES.ACAD_CAREER
                    AND B.STDNT_CAR_NBR = B_ES.STDNT_CAR_NBR
                    AND B.EFFDT = B_ES.EFFDT)
  AND A.EFFDT = B.EFFDT
  AND C.EFFDT = (SELECT MAX(C_ED.EFFDT) FROM SYSADM.PS_ETY_HV1_TBL C_ED
                 WHERE C.INSTITUTION = C_ED.INSTITUTION
                   AND C.ACAD_PROG = C_ED.ACAD_PROG
                   AND C_ED.EFFDT <= SYSDATE)
  AND A.PROG_STATUS = 'AC'
  AND A.EMPLID = :emplid
  `

	// Ejecuta la consulta con el parámetro posicional
	rows, err := Conexion.QueryContext(ctx, query, t)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.ProgramasbyEmplid
	for rows.Next() {
		var item models.ProgramasbyEmplid
		err := rows.Scan(&item.Emplid,
			&item.TipoCarrera,
			&item.Status,
			&item.CodPlanAcademico,
			&item.ProgramaDescripcion,
			&item.ProgramaDuracion)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}
