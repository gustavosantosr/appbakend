package bd

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

/*SendMail end point items*/
func SendMail(code string) {

	m := gomail.NewMessage()

	// Configurar remitente, destinatario y asunto
	m.SetHeader("From", "email@drmonkey.co") // Usa un email verificado en Mandrill
	m.SetHeader("To", "gustavosantosr@gmail.com")
	m.SetHeader("Subject", "Correo de prueba")
	body := fmt.Sprintf("<h1>Hola!</h1><p>El código de seguridad para acceder al sistema es %s.</p>", code)
	m.SetBody("text/html", body)

	// Configurar el servidor SMTP de Mandrill
	//d := gomail.NewDialer("mail.drmonkey.co", 465, "email@drmonkey.co", "Admin4402!")
	//d := gomail.NewDialer("mail.drmonkey.co", 465, "email@drmonkey.co", "Admin4402!")
	//d := gomail.NewDialer("smtp.mandrillapp.com", 587, "email@drmonkey.co", "md-4A6FoxCfhFM7euMC4-WMyw")
	d := gomail.NewDialer("smtp.gmail.com", 587, "poas.consulta.previa@gmail.com", "kkoo uqgi vkhe nzfo ")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Enviar el correo
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error al enviar el correo:", err)
		return
	}
	fmt.Println("Correo enviado con éxito!")
}
