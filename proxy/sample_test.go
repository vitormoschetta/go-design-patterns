package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	name = "John Doe"
	to   = ""
)

func TestSendEmail_WithoutProxy(t *testing.T) {
	// arrange
	emailSender := NewSmtpEmailSender()
	userService := NewUserService(emailSender)

	espectedRes := "Olá " + name + ", seja bem-vindo!"

	// act
	res, err := userService.RegisterUser(name, to)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, espectedRes, res)
}

func TestSendEmail_WithProxy(t *testing.T) {
	// arrange
	emailSender := NewSmtpEmailSender()
	proxyEmailSender := NewProxyEmailSender(emailSender)
	userService := NewUserService(proxyEmailSender)

	espectedRes := ""

	// act
	res, err := userService.RegisterUser(name, to)

	// assert
	assert.NotNil(t, err)
	assert.Equal(t, espectedRes, res)
}
