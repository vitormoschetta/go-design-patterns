package di

import "fmt"

type EmailSender interface {
	SendEmail(to, subject, body string) error
}

type SmtpEmailSender struct{}

func (s *SmtpEmailSender) SendEmail(to, subject, body string) error {
	// Imagine que aqui está a lógica de envio de e-mail para um servidor SMTP
	fmt.Printf("Enviando e-mail para %s com assunto '%s' e corpo '%s'\n", to, subject, body)
	return nil
}

// Defina a estrutura UserService que depende de EmailSender
type UserService struct {
	EmailSender EmailSender
}

func NewUserService(emailSender EmailSender) *UserService {
	return &UserService{EmailSender: emailSender}
}

func (u *UserService) RegisterUser(name, email string) error {
	// Lógica de registro do usuário...
	// Envie um e-mail usando o EmailSender injetado
	return u.EmailSender.SendEmail(email, "Bem-vindo!", "Obrigado por se registrar!")
}
