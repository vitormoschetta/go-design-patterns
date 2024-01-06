package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	name = "John Doe"
	to   = "johndoe@mail.com"
)

func TestSendEmail_WithoutProxy(t *testing.T) {
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

func TestSendEmail_WithProxy(t *testing.T) {
	// arrange
	emailSender := NewSmtpEmailSender()
	proxyEmailSender := NewProxyEmailSender(emailSender)
	userService := NewUserService(proxyEmailSender)

	// act
	res, err := userService.RegisterUser(name, to)

	// assert
	expectedMsg := "Proxy: Olá " + name + ", seja bem-vindo!"
	assert.Nil(t, err)
	assert.Equal(t, expectedMsg, res)
}
