package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteFacade(t *testing.T) {
	// arrange
	facade := NewFacade()

	// act
	err := facade.Execute()

	// assert
	assert.Nil(t, err)
}
