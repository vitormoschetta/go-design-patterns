package di

import "fmt"

type SmtpEmailSender2 struct{}

func (s *SmtpEmailSender2) SendEmail(to, subject, body string) error {
	fmt.Printf("Enviando e-mail para %s com assunto '%s' e corpo '%s'\n", to, subject, body)
	return nil
}

type UserService2 struct{}

func NewUserService2() *UserService2 {
	return &UserService2{}
}

func (u *UserService2) RegisterUser(name, email string) error {
	emailSender := SmtpEmailSender2{} // alto acoplamento
	return emailSender.SendEmail(email, "Bem-vindo!", "Obrigado por se registrar!")
}
