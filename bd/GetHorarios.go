package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/twittor/logger"
	"github.com/gustavosantosr/twittor/models"
)

/*GetHorariosbyEmplid end point items*/
func GetHorariosbyEmplid(t string) ([]*models.Horario, error) {
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
  SELECT DISTINCT 
E.COURSE_TITLE_LONG as CursoNombre,
TO_CHAR(CAST((F.MEETING_TIME_START) AS TIMESTAMP),'HH24.MI.SS.FF') as HoraINICIAL,
TO_CHAR(CAST((F.MEETING_TIME_END) AS TIMESTAMP),'HH24.MI.SS.FF') as HoraFinal,
F.MON as Lunes,
F.TUES as Martes,
F.WED as Miercoles,
F.THURS as Jueves,
F.FRI as Viernes,
F.SAT as Sabado,
F.SUN as Domingo,
TO_CHAR(F.START_DT,'YYYY-MM-DD') as FechaInicial,
TO_CHAR(F.END_DT,'YYYY-MM-DD') as FechaFinal
  FROM SYSADM.PS_ACAD_PROG A, SYSADM.PS_NAMES B, SYSADM.PS_STDNT_ENRL C, SYSADM.PS_CLASS_TBL D, SYSADM.PS_CRSE_CATALOG E, SYSADM.PS_CLASS_MTG_PAT F, SYSADM.PS_TERM_TBL H
  WHERE ( A.EFFDT =
        (SELECT MAX(A_ED.EFFDT) FROM SYSADM.PS_ACAD_PROG A_ED
        WHERE A.EMPLID = A_ED.EMPLID
          AND A.ACAD_CAREER = A_ED.ACAD_CAREER
          AND A.STDNT_CAR_NBR = A_ED.STDNT_CAR_NBR
          AND A_ED.EFFDT <= SYSDATE)
    AND A.EFFSEQ =
        (SELECT MAX(A_ES.EFFSEQ) FROM SYSADM.PS_ACAD_PROG A_ES
        WHERE A.EMPLID = A_ES.EMPLID
          AND A.ACAD_CAREER = A_ES.ACAD_CAREER
          AND A.STDNT_CAR_NBR = A_ES.STDNT_CAR_NBR
          AND A.EFFDT = A_ES.EFFDT)
     AND A.PROG_STATUS = 'AC'
     AND A.EMPLID = B.EMPLID
     AND B.EFFDT =
        (SELECT MAX(B_ED.EFFDT) FROM SYSADM.PS_NAMES B_ED
        WHERE B.EMPLID = B_ED.EMPLID
          AND B.NAME_TYPE = B_ED.NAME_TYPE
          AND B_ED.EFFDT <= SYSDATE)
     AND B.NAME_TYPE = 'PRI'
     AND B.EFF_STATUS = 'A'
     AND A.EMPLID = C.EMPLID
     AND C.INSTITUTION = A.INSTITUTION
     AND C.ACAD_CAREER = D.ACAD_CAREER
     AND C.INSTITUTION = D.INSTITUTION
     AND C.STRM = D.STRM
     AND C.CLASS_NBR = D.CLASS_NBR
     AND D.SESSION_CODE = C.SESSION_CODE
     AND C.STDNT_ENRL_STATUS = 'E'
     AND D.CRSE_ID = E.CRSE_ID
     AND E.EFFDT =
        (SELECT MAX(E_ED.EFFDT) FROM SYSADM.PS_CRSE_CATALOG E_ED
        WHERE E.CRSE_ID = E_ED.CRSE_ID
          AND E_ED.EFFDT <= SYSDATE)
     AND D.CRSE_ID = F.CRSE_ID
     AND D.CRSE_OFFER_NBR = F.CRSE_OFFER_NBR
     AND D.STRM = F.STRM
     AND D.SESSION_CODE = F.SESSION_CODE
     AND D.CLASS_SECTION = F.CLASS_SECTION)
     AND A.EMPLID =:emplid 
  `

	// Ejecuta la consulta con el parámetro posicional
	rows, err := Conexion.QueryContext(ctx, query, t)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Horario
	for rows.Next() {
		var item models.Horario
		err := rows.Scan(&item.CursoNombre,
			&item.HoraInicial,
			&item.HoraFinal,
			&item.Lunes,
			&item.Martes,
			&item.Miercoles,
			&item.Jueves,
			&item.Viernes,
			&item.Sabado,
			&item.Domingo,
			&item.FechaInicial,
			&item.FechaFinal)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}
