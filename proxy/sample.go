package proxy

import "fmt"

// INTERFACE DE ENVIO DE E-MAIL
type IEmailSender interface {
	SendEmail(to, subject, body string) error
}

// IMPLEMENTAÇÃO DE ENVIO DE E-MAIL PARA UM SERVIDOR SMTP
type SmtpEmailSender struct{}

func NewSmtpEmailSender() *SmtpEmailSender {
	return &SmtpEmailSender{}
}

func (s *SmtpEmailSender) SendEmail(to, subject, body string) error {
	fmt.Printf("Enviando e-mail para %s com assunto '%s' e corpo '%s'\n", to, subject, body)
	return nil
}

// PROXY/INTERCEPTOR QUE ADICIONA UM COMPORTAMENTO AO ENVIO DE E-MAIL
type ProxyEmailSender struct { // Meu proxy pode ser ou não uma interface, depende da flexibilidade que eu quero ter com ele.
	// Como nesse caso é um serviço de logging e validação internos, não preciso de flexibilidade, então posso usar uma struct
	EmailSender IEmailSender
}

func NewProxyEmailSender(emailSender IEmailSender) *ProxyEmailSender {
	return &ProxyEmailSender{
		EmailSender: emailSender,
	}
}

func (p *ProxyEmailSender) SendEmail(to, subject, body string) error {
	// Agregamos um objeto intermediario que vai fazer a chamada do metodo SendEmail do objeto EmailSender
	// A ideia é tratar o dado e fazer logging ou alguma validação antes de chamar o metodo SendEmail do objeto EmailSender
	fmt.Println("Adicionando comportamento")
	fmt.Printf("Enviando e-mail para %s com assunto '%s' e corpo '%s'\n", to, subject, body)
	return p.EmailSender.SendEmail(to, subject, body)
}

type UserService struct {
	EmailSender IEmailSender
}

func NewUserService(emailSender IEmailSender) *UserService {
	return &UserService{
		EmailSender: emailSender,
	}
}

func (u *UserService) RegisterUser(name, email string) error {
	return u.EmailSender.SendEmail(email, "Bem-vindo", "Olá "+name+", seja bem-vindo!")
}
