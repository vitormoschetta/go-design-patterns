package proxy

import "fmt"

type EmailSender interface {
	SendEmail(to, subject, body string) error
}

type SmtpEmailSender struct{}

func (s *SmtpEmailSender) SendEmail(to, subject, body string) error {
	// Imagine que aqui está a lógica de envio de e-mail para um servidor SMTP
	return nil
}

// Meu proxy pode ser ou não uma interface, depende da flexibilidade que eu quero ter com ele.
// Como nesse caso é um serviço de logging e validação internos, não preciso de flexibilidade, então posso usar uma struct
type ProxyEmailSender struct {
	EmailSender EmailSender
}

func NewProxyEmailSender(emailSender EmailSender) *ProxyEmailSender {
	return &ProxyEmailSender{
		EmailSender: emailSender,
	}
}

func (p *ProxyEmailSender) SendEmail(to, subject, body string) error {
	// Agregamos um objeto intermediario que vai fazer a chamada do metodo SendEmail do objeto EmailSender
	// A ideia é tratar o dado e fazer logging ou alguma validação antes de chamar o metodo SendEmail do objeto EmailSender
	fmt.Printf("Enviando e-mail para %s com assunto '%s' e corpo '%s'\n", to, subject, body)
	return p.EmailSender.SendEmail(to, subject, body)
}

type UserService struct {
	ProxyEmailSender ProxyEmailSender
}

func NewUserService(proxyEmailSender ProxyEmailSender) *UserService {
	return &UserService{
		ProxyEmailSender: proxyEmailSender,
	}
}

func (u *UserService) RegisterUser(name, email string) error {
	return u.ProxyEmailSender.SendEmail(email, "Bem-vindo", "Olá "+name+", seja bem-vindo!")
}
