package bd

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

/*SendMail end point items*/
func SendMail() {

	m := gomail.NewMessage()

	// Configurar remitente, destinatario y asunto
	m.SetHeader("From", "tuemail@tudominio.com") // Usa un email verificado en Mandrill
	m.SetHeader("To", "gustavosantosr@gmail.com")
	m.SetHeader("Subject", "Correo de prueba")
	m.SetBody("text/html", "<h1>Hola!</h1><p>El codigo de seguridad para acceder al sistema es 111111g.</p>")

	// Configurar el servidor SMTP de Mandrill
	d := gomail.NewDialer("mail.drmonkey.co", 465, "email@drmonkey.co", "Admin4402!")

	// Enviar el correo
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error al enviar el correo:", err)
		return
	}
	fmt.Println("Correo enviado con éxito!")
}
