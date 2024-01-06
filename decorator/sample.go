package decorator

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

// SERVIÇO DE USUÁRIO QUE UTILIZA O ENVIO DE E-MAIL
type UserService struct {
	EmailSender IEmailSender
}

func NewUserService(emailSender IEmailSender) *UserService {
	return &UserService{EmailSender: emailSender}
}

func (u *UserService) RegisterUser(name, email string) (string, error) {
	body := "Olá " + name + ", seja bem-vindo!"
	return u.EmailSender.SendEmail(email, "Bem-vindo", body)
}

// DECORATOR QUE ADICIONA UM COMPORTAMENTO AO REGISTRO DE USUÁRIO
type IDecorator interface {
	IEmailSender // Herda os contratos da interface IEmailSender
	AddedBehavior() string
}

type DecoratorEmailSender struct {
	EmailSender IEmailSender
}

func (d *DecoratorEmailSender) AddedBehavior() string {
	return "Comportamento adicionado"
}

func (d *DecoratorEmailSender) SendEmail(to, subject, body string) (string, error) {
	return d.EmailSender.SendEmail(to, subject, body)
}

// SERVIÇO DE USUÁRIO QUE UTILIZA O DECORATOR
type UserService2 struct {
	Decorator IDecorator
}

func NewUserService2(decorator IDecorator) *UserService2 {
	return &UserService2{Decorator: decorator}
}

func (u *UserService2) RegisterUser(name, email string) (string, error) {
	res := u.Decorator.AddedBehavior() // Comportamento adicionado
	body := "Olá " + name + ", seja bem-vindo!"
	return u.Decorator.SendEmail(email, "Bem-vindo", body+" "+res)
}
