package proxy

import "fmt"

// Defina a estrutura UserService que depende de EmailSender
type UserService2 struct {
	EmailSender IEmailSender
}

func NewUserService2(emailSender IEmailSender) *UserService2 {
	return &UserService2{
		EmailSender: emailSender,
	}
}

func (u *UserService2) RegisterUser(name, email string) error {
	// Lógica de registro do usuário...
	// Como não temos um proxy para fazer o logging e validação, vamos fazer isso aqui mesmo:
	fmt.Printf("Registrando usuário %s com e-mail %s\n", name, email)
	// Envie um e-mail usando o EmailSender injetado
	return u.EmailSender.SendEmail(email, "Bem-vindo!", "Obrigado por se registrar!")
}
