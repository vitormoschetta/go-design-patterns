package di

import "fmt"

// INTERFACE DE ENVIO DE E-MAIL
type IEmailSender interface {
	SendEmail(to, subject, body string) error
}

// IMPLEMENTAÇÃO DE ENVIO DE E-MAIL PARA UM SERVIDOR SMTP
type SmtpEmailSender struct{}

func (s *SmtpEmailSender) SendEmail(to, subject, body string) error {
	fmt.Printf("Enviando e-mail para %s com assunto '%s' e corpo '%s'\n", to, subject, body)
	return nil
}

// SERVIÇO DE USUÁRIO QUE UTILIZA O ENVIO DE E-MAIL
type UserService struct {
	EmailSender IEmailSender
}

func NewUserService(emailSender IEmailSender) *UserService {
	return &UserService{EmailSender: emailSender}
}

func (u *UserService) RegisterUser(name, email string) error {
	// Lógica de registro do usuário...
	// Envie um e-mail usando o EmailSender injetado
	return u.EmailSender.SendEmail(email, "Bem-vindo!", "Obrigado por se registrar!")
}
