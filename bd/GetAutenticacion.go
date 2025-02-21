package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gustavosantosr/twittor/logger"
	"github.com/gustavosantosr/twittor/models"
)

/*GetAutenticacion end point items*/
func GetAutenticacion() ([]*models.Autenticacion, error) {
	err := Conexion.Ping()
	if err != nil {
		logger.WriteLogger(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		logger.WriteLogger(fmt.Sprintf("%+v", err.Error()))
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `
  SELECT * 
  FROM PS_UCA_AUTHENTICATION_APP
  `

	// Ejecuta la consulta con el parámetro posicional
	rows, err := Conexion.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.Autenticacion
	for rows.Next() {
		var item models.Autenticacion
		err := rows.Scan(&item.DOCUMENT,
			&item.CODE,
			&item.REGISTRATION_DATE,
			&item.STATE)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}
