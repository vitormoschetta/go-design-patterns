package decorator

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
	return u.EmailSender.SendEmail(email, "Bem-vindo!", "Obrigado por se registrar!")
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

func (d *DecoratorEmailSender) SendEmail(to, subject, body string) error {
	fmt.Println("Adicionando comportamento")
	return d.EmailSender.SendEmail(to, subject, body)
}

// SERVIÇO DE USUÁRIO QUE UTILIZA O DECORATOR
type UserService2 struct {
	Decorator IDecorator
}

func NewUserService2(decorator IDecorator) *UserService2 {
	return &UserService2{Decorator: decorator}
}

func (u *UserService2) RegisterUser(name, email string) error {
	u.Decorator.AddedBehavior() // Comportamento adicionado
	return u.Decorator.SendEmail(email, "Bem-vindo!", "Obrigado por se registrar!")
}
