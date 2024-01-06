package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockEmailSender struct {
	SendEmailFunc func(to, subject, body string) error
}

func (m *MockEmailSender) SendEmail(to, subject, body string) error {
	return m.SendEmailFunc(to, subject, body)
}

func TestSendEmail_WithoutDecorator(t *testing.T) {
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

func TestSendEmail_WithDecorator(t *testing.T) {
	// arrange
	mockEmailSender := &MockEmailSender{
		SendEmailFunc: func(to, subject, body string) error {
			return nil
		},
	}
	decorator := &DecoratorEmailSender{EmailSender: mockEmailSender}
	userService := NewUserService2(decorator)

	// act
	err := userService.RegisterUser("John Doe", "user@mail.com")

	// assert
	assert.Nil(t, err)
}
