package proxy

import "fmt"

// INTERFACE DE ENVIO DE E-MAIL
type IEmailSender interface {
	SendEmail(to, subject, body string) (string, error)
}

// IMPLEMENTAÇÃO DE ENVIO DE E-MAIL
type SmtpEmailSender struct{}

func NewSmtpEmailSender() *SmtpEmailSender {
	return &SmtpEmailSender{}
}

func (s *SmtpEmailSender) SendEmail(to, subject, body string) (string, error) {
	msg := fmt.Sprintf("Enviando e-mail para %s, assunto: %s, corpo: %s", to, subject, body)
	fmt.Println(msg)
	return body, nil
}

// PROXY/INTERCEPTOR QUE IMPLEMENTA A INTERFACE DE ENVIO DE E-MAIL E FUNCIONA COMO UM OBJETO INTERMEDIARIO ENTRE O USUARIO E O OBJETO REAL DE ENVIO DE E-MAIL
type ProxyEmailSender struct {
	EmailSender IEmailSender
}

func NewProxyEmailSender(emailSender IEmailSender) *ProxyEmailSender {
	return &ProxyEmailSender{
		EmailSender: emailSender,
	}
}

func (p *ProxyEmailSender) SendEmail(to, subject, body string) (string, error) {
	// Aqui podemos adicionar comportamentos extras, seja antes ou depois de chamar o metodo SendEmail do objeto real, funcionando como um interceptor/middleware

	// Chama o metodo SendEmail do objeto real
	msg, err := p.EmailSender.SendEmail(to, subject, body)
	if err != nil {
		return "", err
	}
	msg = "Proxy: " + msg
	return msg, nil
}

type UserService struct {
	EmailSender IEmailSender
}

func NewUserService(emailSender IEmailSender) *UserService {
	return &UserService{
		EmailSender: emailSender,
	}
}

func (u *UserService) RegisterUser(name, email string) (string, error) {
	body := "Olá " + name + ", seja bem-vindo!"
	return u.EmailSender.SendEmail(email, "Bem-vindo", body)
}
