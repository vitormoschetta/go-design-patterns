package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail2_Ok(t *testing.T) {
	// arrange
	mockEmailSender := &MockEmailSender{
		SendEmailFunc: func(to, subject, body string) error {
			return nil
		},
	}
	userService := NewUserService2(mockEmailSender)

	// act
	err := userService.RegisterUser("John Doe", "user@mail.com")

	// assert
	assert.Nil(t, err)
}
