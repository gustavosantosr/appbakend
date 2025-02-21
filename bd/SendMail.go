package bd

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

/*SendMail end point items*/
func SendMail() {

	m := gomail.NewMessage()

	// Configuración del correo
	m.SetHeader("From", "gustavosantosr@gmail.com")
	m.SetHeader("To", "gustavosantosr@gmail.com")
	m.SetHeader("Subject", "Correo de prueba")
	m.SetBody("text/html", "<h1>Hola!</h1><p>Este es un correo enviado desde Golang.</p>")

	// Configuración del servidor SMTP
	d := gomail.NewDialer("smtp.gmail.com", 587, "gustavosantosr@gmail.com", "Gust2910!*")

	// Enviar el correo
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error al enviar el correo:", err)
		return
	}
	fmt.Println("Correo enviado con éxito!")
}
