package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail2_Ok(t *testing.T) {
	// arrange
	// Aqui não temos a opção de injetar o serviço de email que queremos utilizar, ele é criado dentro do UserService2,
	// e por isso não temos como testar o UserService2 com um mock de email. Dessa forma só poderíamos testar o UserService2 com um serviço de email real.
	// Ou seja, seria um teste de integração e não de unidade. E também não seríamos capaz de alterar o provedor de email sem alterar o código do UserService2.
	userService := NewUserService2()

	// act
	err := userService.RegisterUser("John Doe", "user@mail.com")

	// assert
	assert.Nil(t, err)
}
