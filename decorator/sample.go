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
type IUserService interface {
	RegisterUser(name, email string) (string, error)
}

type UserService struct {
	EmailSender IEmailSender
}

func NewUserService(emailSender IEmailSender) IUserService {
	return &UserService{EmailSender: emailSender}
}

func (u *UserService) RegisterUser(name, email string) (string, error) {
	body := "Olá " + name + ", seja bem-vindo!"
	return u.EmailSender.SendEmail(email, "Bem-vindo", body)
}

// DECORATOR DE USER SERVICE
type UserServiceDecorator struct {
	UserService IUserService
}

func NewUserServiceDecorator(userService IUserService) IUserService {
	return &UserServiceDecorator{UserService: userService}
}

func (u *UserServiceDecorator) RegisterUser(name, email string) (string, error) {
	msg, _ := u.AddBehavior()
	fmt.Println(msg)
	res, err := u.UserService.RegisterUser(name, email)
	return res + msg, err
}

func (u *UserServiceDecorator) AddBehavior() (string, error) {
	msg := " Comportamento adicionado"
	fmt.Println(msg)
	return msg, nil
}
