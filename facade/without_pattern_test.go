package facade

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteFacade2(t *testing.T) {
	// arrange
	system1 := &System1{}
	system2 := &System2{}

	// act
	result1, err1 := system1.Execute("param1")
	fmt.Println(result1)

	result2, err2 := system2.Execute("param2")
	fmt.Println(result2)

	// assert
	assert.Nil(t, err1)
	assert.Nil(t, err2)
}
