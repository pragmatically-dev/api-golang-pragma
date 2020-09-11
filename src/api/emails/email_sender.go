package emails

import (
	"bytes"
	"fmt"
	"log"
	"mime/quotedprintable"
	"net/smtp"

	"github.com/pragmatically-dev/apirest/src/config"
)

//Sender estructura definida para centralizar los datos necesarios para enviar el email
type Sender struct {
	HostURL      string
	HostPort     string
	EmailSender  string
	Password     string
	EmailReciver string
}

//SendEmail se encarga de usar un cliente smtp para enviar mails a diferentes receptores
func (s *Sender) SendEmail(msg []byte, reciver string) (bool, error) {

	auth := smtp.PlainAuth(
		"",
		s.EmailSender,
		s.Password,
		s.HostURL,
	)

	err := smtp.SendMail(
		s.HostURL+":"+s.HostPort,
		auth,
		s.EmailSender,
		[]string{s.EmailReciver},
		msg,
	)
	if err != nil {
		log.Println("no se pudo enviar el email")
		return false, err
	}
	return true, nil
}

//Test Funciona para testear el envio de mensajes
func Test() {
	//TEST ENVIO DE MAILS
	sender := Sender{
		EmailSender:  "pragmatically.dev@gmail.com",
		Password:     config.EPASS,
		EmailReciver: "totonieva3@gmail.com",
		HostURL:      "smtp.gmail.com",
		HostPort:     "587",
	}

	header := make(map[string]string)
	header["From"] = sender.EmailSender
	header["To"] = sender.EmailReciver
	header["Subject"] = "Probando la libreria smtp de go"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", "text/html")
	header["Content-Disposition"] = "inline"
	header["Content-Transfer-Encoding"] = "quoted-printable"

	headerMessage := "Test"
	for key, value := range header {
		headerMessage += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	body := `
		<h1>Este es un email de prueba enviado desde el servidor :,)</h1>
		<h1>Tu codigo de verificacion es: 9000  \n\n </h1>
		`

	var bodyMessage bytes.Buffer
	/*
		Quoted-printable, o codificación QP, es una codificación que usa caracteres imprimibles
		para transmitir datos de 8 bit sobre un protocolo que solamente soporta 7 bit.
		Está definido como un content transfer encoding de MIME para ser usado en mensajes de
		correo electrónico de Internet
	*/
	temp := quotedprintable.NewWriter(&bodyMessage)
	temp.Write([]byte(body))
	temp.Close()

	finalMessage := headerMessage + "\r\n" + bodyMessage.String()

	iSended, err := sender.SendEmail([]byte(finalMessage), sender.EmailReciver)
	if err != nil {
		log.Println(err)
	}
	if iSended {
		log.Println("el mensaje fue enviado")
	}
	//END ENVIO DE MAILS
}
