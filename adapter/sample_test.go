package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPrinterService(t *testing.T) {
	// arrange
	legacyPrinter := &LegacyPrinter{}
	adapter := &PrinterAdapter{LegacyPrinter: legacyPrinter}
	service := NewPrinterService(adapter)

	// act
	msg, err := service.Print("Hello World")

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "Legacy Printer: Hello World", msg)

}
