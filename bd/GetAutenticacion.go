package bd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gustavosantosr/twittor/logger"
	"github.com/gustavosantosr/twittor/models"
)

type AutenticacionEmail struct {
	Email string
	Code  string
}
type AutenticacionUsuario struct {
	Email string
}

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
	//InsertAutenticacion()
	//SendMail()

	query := `
  SELECT * 
  FROM SYSADM.PS_UCA_AUTHENTICATION_APP
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

/*GetAutenticacionEmail end point items*/
func GetAutenticacionEmail() ([]*models.AutenticacionEmail, error) {
	err := Conexion.Ping()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error al ejecutar la consulta: %+v", err.Error()))
		log.Fatal(fmt.Sprintf("%+v", err.Error()))
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `
    SELECT 
    A.DOCUMENT, 
    A.CODE,
	A.REGISTRATION_DATE,   
    A.STATE,
    C.EMAIL_ADDR
FROM SYSADM.PS_UCA_AUTHENTICATION_APP A
LEFT JOIN SYSADM.PS_PERS_NID B 
    ON B.NATIONAL_ID = A.DOCUMENT
LEFT JOIN SYSADM.PS_EMAIL_ADDRESSES C 
    ON C.EMPLID = B.EMPLID 
WHERE 
    C.PREF_EMAIL_FLAG = 'Y' 
    AND B.PRIMARY_NID = 'Y'
  `

	// Ejecuta la consulta con el parámetro posicional
	rows, err := Conexion.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(fmt.Sprintf("%+v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var resultados []*models.AutenticacionEmail
	for rows.Next() {
		var item models.AutenticacionEmail
		err := rows.Scan(&item.DOCUMENT,
			&item.CODE,
			&item.REGISTRATION_DATE,
			&item.STATE,
			&item.EMAIL)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return resultados, err
		}
		resultados = append(resultados, &item)

	}
	return resultados, nil
}

// GetEmailAndCode obtiene el correo y el código basado en el documento
func GetEmailAndCode(document string) (*AutenticacionEmail, error) {
	// Verificar conexión a la base de datos
	err := Conexion.Ping()
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
		return nil, err
	}
	InsertAutenticacion(document)
	// Definir contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Query SQL con un parámetro
	query := `
		SELECT 
		A.CODE,
		C.EMAIL_ADDR
	FROM SYSADM.PS_UCA_AUTHENTICATION_APP A
	LEFT JOIN SYSADM.PS_PERS_NID B 
		ON B.NATIONAL_ID = A.DOCUMENT
	LEFT JOIN SYSADM.PS_EMAIL_ADDRESSES C 
		ON C.EMPLID = B.EMPLID 
	WHERE 
		A.DOCUMENT = :1
		AND C.PREF_EMAIL_FLAG = 'Y' 
		AND B.PRIMARY_NID = 'Y'
	ORDER BY A.REGISTRATION_DATE DESC
	FETCH FIRST 1 ROW ONLY`

	// Variables para almacenar los valores retornados
	var result AutenticacionEmail

	// Ejecutar la consulta con el documento como parámetro
	err = Conexion.QueryRowContext(ctx, query, document).Scan(&result.Code, &result.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontró un resultado
			return nil, fmt.Errorf("no se encontraron datos para el documento %s", document)
		}
		log.Printf("Error ejecutando la consulta: %v", err)
		return nil, err
	}
	log.Printf("codigo: %v", result.Code)
	SendMail(result.Code)
	return &result, nil
}

// GetEmailByCredentials obtiene el correo electrónico basado en usuario y contraseña.
func GetEmailByCredentials(code, document string) (*AutenticacionUsuario, error) {
	// Verificar conexión a la base de datos
	err := Conexion.Ping()
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
		return nil, err
	}

	// Definir contexto con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Query SQL para obtener el correo electrónico basado en usuario y contraseña
	query := `
		SELECT 
		A.CODE,
		C.EMAIL_ADDR
	FROM SYSADM.PS_UCA_AUTHENTICATION_APP A
	LEFT JOIN SYSADM.PS_PERS_NID B 
		ON B.NATIONAL_ID = A.DOCUMENT
	LEFT JOIN SYSADM.PS_EMAIL_ADDRESSES C 
		ON C.EMPLID = B.EMPLID 
	WHERE 
	A.CODE = :1
		AND A.DOCUMENT = :2
		
		AND C.PREF_EMAIL_FLAG = 'Y' 
		AND B.PRIMARY_NID = 'Y'
	ORDER BY A.REGISTRATION_DATE DESC
	FETCH FIRST 1 ROW ONLY`

	// Variable para almacenar el correo electrónico
	var result AutenticacionUsuario

	// Ejecutar la consulta
	err = Conexion.QueryRowContext(ctx, query, code, document).Scan(&result.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontró un resultado
			return nil, fmt.Errorf("usuario o contraseña incorrectos")
		}
		log.Printf("Error ejecutando la consulta: %v", err)
		return nil, err
	}

	log.Printf("Correo encontrado: %v", result.Email)
	return &result, nil
}
