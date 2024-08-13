package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/twittor/logger"
	"github.com/gustavosantosr/twittor/models"
)

/*GetMaterias end point items*/
func GetMaterias() ([]*models.Materias, error) {
	err := Conexion.Ping()
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		logger.WriteLogger(fmt.Sprintf("%+v", err.Error()))
		return nil, err
	}
	query := `SELECT 
    DISTINCT A.EMPLID, 
    A.ACAD_CAREER as TipoCarrera, 
    A.ACAD_PROG as CodAcademicoPrograma, 
    B.ACAD_PLAN as CodAcademicoPlan, 
    C.ETY_PROG_DESCR_LG as DescripcionPrograma, 
    G.CRSE_ID as CursoId, 
    G.STRM as Strm, 
    G.CATALOG_NBR AS CatalogoNombre, 
    I.COURSE_TITLE_LONG as CursoNombre, 
    H.DESCRSHORT as CodigoSemestre
FROM 
    SYSADM.PS_ACAD_PROG A
JOIN 
    SYSADM.PS_ACAD_PLAN B ON A.EMPLID = B.EMPLID
                    AND A.ACAD_CAREER = B.ACAD_CAREER
                    AND A.STDNT_CAR_NBR = B.STDNT_CAR_NBR
JOIN 
    SYSADM.PS_ETY_HV1_TBL C ON C.INSTITUTION = A.INSTITUTION
                      AND C.ACAD_PROG = A.ACAD_PROG
JOIN 
    SYSADM.PS_STDNT_ENRL D ON A.EMPLID = D.EMPLID
                     AND A.ACAD_CAREER = D.ACAD_CAREER
                     AND D.INSTITUTION = A.INSTITUTION
                     AND D.STDNT_ENRL_STATUS = 'E'
JOIN 
    SYSADM.PS_CLST_MAIN_TBL E ON A.ACAD_CAREER = E.ACAD_CAREER
                     AND A.ACAD_PROG = E.ACAD_PROG --Comentar esta línea para sacar total sin verificar contra plan de estudios
                     AND B.ACAD_PLAN = E.ACAD_PLAN --Comentar esta línea para sacar total sin verificar contra plan de estudios
JOIN 
    SYSADM.PS_CLST_DETL_TBL F ON E.COURSE_LIST = F.COURSE_LIST
JOIN 
    SYSADM.PS_CLASS_TBL G ON D.ACAD_CAREER = G.ACAD_CAREER
                   AND D.INSTITUTION = G.INSTITUTION
                   AND D.STRM = G.STRM
                   AND D.CLASS_NBR = G.CLASS_NBR
                   AND G.SESSION_CODE = D.SESSION_CODE
                   AND G.CRSE_ID = F.CRSE_ID
JOIN 
    SYSADM.PS_TERM_TBL H ON G.STRM = H.STRM
                  AND H.INSTITUTION = G.INSTITUTION
                  AND H.ACAD_CAREER = G.ACAD_CAREER
JOIN 
    SYSADM.PS_CRSE_CATALOG I ON G.CRSE_ID = I.CRSE_ID
WHERE 
    A.EFFDT = (SELECT MAX(A_ED.EFFDT) FROM SYSADM.PS_ACAD_PROG A_ED
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
    AND C.EFFDT = (SELECT MAX(C_ED.EFFDT) FROM SYSADM.PS_ETY_HV1_TBL C_ED
                      WHERE C.INSTITUTION = C_ED.INSTITUTION
                        AND C.ACAD_PROG = C_ED.ACAD_PROG
                        AND C_ED.EFFDT <= SYSDATE)
    AND E.EFFDT = (SELECT MAX(E_ED.EFFDT) FROM SYSADM.PS_CLST_MAIN_TBL E_ED
                      WHERE E.COURSE_LIST = E_ED.COURSE_LIST
                        AND E_ED.EFFDT <= SYSDATE)
    AND F.EFFDT = (SELECT MAX(F_ED.EFFDT) FROM SYSADM.PS_CLST_DETL_TBL F_ED
                      WHERE F.COURSE_LIST = F_ED.COURSE_LIST
                        AND F_ED.EFFDT <= SYSDATE)
    --AND D.STRM IN ('1830')-- semestre
    AND D.GRADING_BASIS_ENRL<>'NON'
    AND I.EFFDT = (SELECT MAX(I_ED.EFFDT) FROM SYSADM.PS_CRSE_CATALOG I_ED
                      WHERE I.CRSE_ID = I_ED.CRSE_ID
                        AND I_ED.EFFDT <= SYSDATE)
   -- AND A.ACAD_CAREER = 'TECN'--grado academico
    --AND A.PROG_STATUS = 'AC'
  --  AND A.ACAD_PROG IN ('237',
--'255',
--'292',
--'372',
--'373')--codigos de programa
    AND A.EMPLID IN ('0000052403')`
	logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", query))
	rows, err := Conexion.Query(query)
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Materias
	for rows.Next() {
		var item models.Materias
		err := rows.Scan(&item.EMPLID,
			&item.TipoCarrera,
			&item.CodAcademicoPrograma,
			&item.CodAcademicoPlan,
			&item.DescripcionPrograma,
			&item.CursoId,
			&item.Strm,
			&item.CatalogoNombre,
			&item.CursoNombre,
			&item.CodigoSemestre)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

func GetMateriasbyEmplid(emplid string, codprog string) ([]*models.Semestre, error) {
	err := Conexion.Ping()
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		logger.WriteLogger(fmt.Sprintf("%+v", err.Error()))
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	fmt.Println("parametro ts\t", emplid)
	query := `
	SELECT 
	    DISTINCT  
	    A.ACAD_CAREER as TipoCarrera, 
	    A.ACAD_PROG as CodAcademicoPrograma, 
	    B.ACAD_PLAN as CodAcademicoPlan, 
	    C.ETY_PROG_DESCR_LG as DescripcionPrograma, 
	    G.CRSE_ID as CursoId, 
	    G.STRM as Strm, 
	    G.CATALOG_NBR AS CatalogoNombre, 
	    I.COURSE_TITLE_LONG as CursoNombre,
		D.CRSE_GRADE_INPUT as NotaDefinitiva, 
	    H.DESCRSHORT as CodigoSemestre
	FROM 
	    SYSADM.PS_ACAD_PROG A
	JOIN 
	    SYSADM.PS_ACAD_PLAN B ON A.EMPLID = B.EMPLID
	                    AND A.ACAD_CAREER = B.ACAD_CAREER
	                    AND A.STDNT_CAR_NBR = B.STDNT_CAR_NBR
	JOIN 
	    SYSADM.PS_ETY_HV1_TBL C ON C.INSTITUTION = A.INSTITUTION
	                      AND C.ACAD_PROG = A.ACAD_PROG
	JOIN 
	    SYSADM.PS_STDNT_ENRL D ON A.EMPLID = D.EMPLID
	                     AND A.ACAD_CAREER = D.ACAD_CAREER
	                     AND D.INSTITUTION = A.INSTITUTION
	                     AND D.STDNT_ENRL_STATUS = 'E'
	JOIN 
	    SYSADM.PS_CLST_MAIN_TBL E ON A.ACAD_CAREER = E.ACAD_CAREER
	                     AND A.ACAD_PROG = E.ACAD_PROG
	                     AND B.ACAD_PLAN = E.ACAD_PLAN
	JOIN 
	    SYSADM.PS_CLST_DETL_TBL F ON E.COURSE_LIST = F.COURSE_LIST
	JOIN 
	    SYSADM.PS_CLASS_TBL G ON D.ACAD_CAREER = G.ACAD_CAREER
	                   AND D.INSTITUTION = G.INSTITUTION
	                   AND D.STRM = G.STRM
	                   AND D.CLASS_NBR = G.CLASS_NBR
	                   AND G.SESSION_CODE = D.SESSION_CODE
	                   AND G.CRSE_ID = F.CRSE_ID
	JOIN 
	    SYSADM.PS_TERM_TBL H ON G.STRM = H.STRM
	                  AND H.INSTITUTION = G.INSTITUTION
	                  AND H.ACAD_CAREER = G.ACAD_CAREER
	JOIN 
	    SYSADM.PS_CRSE_CATALOG I ON G.CRSE_ID = I.CRSE_ID
	WHERE 
	    A.EFFDT = (SELECT MAX(A_ED.EFFDT) FROM SYSADM.PS_ACAD_PROG A_ED
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
	    AND C.EFFDT = (SELECT MAX(C_ED.EFFDT) FROM SYSADM.PS_ETY_HV1_TBL C_ED
	                      WHERE C.INSTITUTION = C_ED.INSTITUTION
	                        AND C.ACAD_PROG = C_ED.ACAD_PROG
	                        AND C_ED.EFFDT <= SYSDATE)
	    AND E.EFFDT = (SELECT MAX(E_ED.EFFDT) FROM SYSADM.PS_CLST_MAIN_TBL E_ED
	                      WHERE E.COURSE_LIST = E_ED.COURSE_LIST
	                        AND E_ED.EFFDT <= SYSDATE)
	    AND F.EFFDT = (SELECT MAX(F_ED.EFFDT) FROM SYSADM.PS_CLST_DETL_TBL F_ED
	                      WHERE F.COURSE_LIST = F_ED.COURSE_LIST
	                        AND F_ED.EFFDT <= SYSDATE)
	    AND D.GRADING_BASIS_ENRL <> 'NON'
	    AND I.EFFDT = (SELECT MAX(I_ED.EFFDT) FROM SYSADM.PS_CRSE_CATALOG I_ED
	                      WHERE I.CRSE_ID = I_ED.CRSE_ID
	                        AND I_ED.EFFDT <= SYSDATE)
	    AND A.EMPLID = :1
		AND A.ACAD_PROG = :2
		order by H.DESCRSHORT desc`

	logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", query))
	rows, err := Conexion.QueryContext(ctx, query, emplid, codprog)
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	semestresMap := make(map[string]*models.Semestre)

	for rows.Next() {
		var materia models.Materias
		err := rows.Scan(&materia.TipoCarrera, &materia.CodAcademicoPrograma, &materia.CodAcademicoPlan,
			&materia.DescripcionPrograma, &materia.CursoId, &materia.Strm, &materia.CatalogoNombre,
			&materia.CursoNombre, &materia.NotaDefinitiva, &materia.CodigoSemestre)
		if err != nil {
			logger.WriteLogger(fmt.Sprintf("Error al leer las filas: %+v", err.Error()))
			return nil, err
		}

		// Verifica si el semestre ya existe en el mapa
		if semestre, exists := semestresMap[materia.CodigoSemestre]; exists {
			// Agrega la materia al semestre existente
			semestre.Materias = append(semestre.Materias, &materia)
		} else {
			// Crea un nuevo semestre y agrega la materia
			semestresMap[materia.CodigoSemestre] = &models.Semestre{
				Numero:   materia.CodigoSemestre,
				Materias: []*models.Materias{&materia},
			}
		}
	}
	// Convertir el mapa en un slice de semestres
	var semestres []*models.Semestre
	for _, semestre := range semestresMap {
		semestres = append(semestres, semestre)
	}

	return semestres, nil

}
