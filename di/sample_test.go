package di

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockEmailSender struct {
	SendEmailFunc func(to, subject, body string) error
}

func (m *MockEmailSender) SendEmail(to, subject, body string) error {
	return m.SendEmailFunc(to, subject, body)
}

func TestSendEmail_Ok(t *testing.T) {
	// arrange
	mockEmailSender := &MockEmailSender{
		SendEmailFunc: func(to, subject, body string) error {
			return nil
		},
	}
	userService := NewUserService(mockEmailSender)

	// act
	err := userService.RegisterUser("John Doe", "user@mail.com")

	// assert
	assert.Nil(t, err)
}

func TestSendEmail_Error(t *testing.T) {
	// arrange
	mockEmailSender := &MockEmailSender{
		SendEmailFunc: func(to, subject, body string) error {
			return errors.New("error sending email")
		},
	}
	userService := NewUserService(mockEmailSender)

	// act
	err := userService.RegisterUser("John Doe", "user@mail.com")

	// assert
	assert.NotNil(t, err)
}
