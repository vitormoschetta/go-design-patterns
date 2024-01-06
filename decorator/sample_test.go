package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	name = "John Doe"
	to   = "johndoe@mail.com"
)

func TestSendEmail_WithoutDecorator(t *testing.T) {
	// arrange
	emailSender := NewSmtpEmailSender()
	userService := NewUserService(emailSender)

	// act
	res, err := userService.RegisterUser(name, to)

	// assert
	espectedRes := "Olá " + name + ", seja bem-vindo!"
	assert.Nil(t, err)
	assert.Equal(t, espectedRes, res)
}

func TestSendEmail_WithDecorator(t *testing.T) {
	// arrange
	emailSender := NewSmtpEmailSender()
	decorator := &DecoratorEmailSender{EmailSender: emailSender}
	userService := NewUserService2(decorator)

	// act
	res, err := userService.RegisterUser(name, to)

	// assert
	expectedMsg := "Olá " + name + ", seja bem-vindo! Comportamento adicionado"
	assert.Nil(t, err)
	assert.Equal(t, expectedMsg, res)
}
