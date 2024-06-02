package bd

import (
	"fmt"

	"github.com/gustavosantosr/twittor/logger"
	"github.com/gustavosantosr/twittor/models"
)

/*GetConsolidados() realiza la busqueda de consolidados diferentes en la base de datos*/
func GetConsolidados() ([]*models.Consolidados, error) {
	// Verificar la conexión a la base de datos
	err := Conexion.Ping()
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		logger.WriteLogger(fmt.Sprintf("%+v", err.Error()))
		return nil, err
	}

	query := "SELECT EMPLID FROM SYSADM.PS_PERS_NID"
	rows, err := Conexion.Query(query)
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Consolidados
	for rows.Next() {
		var item1 models.Consolidados
		err := rows.Scan(&item1.EMPLID)
		if err != nil {
			logger.WriteLogger(fmt.Sprintf("Error al leer las filas: %+v", err.Error()))
			return nil, err
		}
		resultados = append(resultados, &item1)
	}

	return resultados, nil
}
