package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail_WithoutProxy(t *testing.T) {
	// arrange
	emailSender := NewSmtpEmailSender()
	userService := NewUserService(emailSender)

	// act
	err := userService.RegisterUser("John Doe", "user@mail.com")

	// assert
	assert.Nil(t, err)
}

func TestSendEmail_WithProxy(t *testing.T) {
	// arrange
	emailSender := NewSmtpEmailSender()
	proxyEmailSender := NewProxyEmailSender(emailSender)
	userService := NewUserService(proxyEmailSender)

	// act
	err := userService.RegisterUser("John Doe", "user@mail.com")

	// assert
	assert.Nil(t, err)
}
