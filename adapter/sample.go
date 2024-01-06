package adapter

import "fmt"

// INTERFACE DESEJADA
type IPrinter interface {
	Print(message string) (string, error)
}

// INTERFACE LEGADA
type LegacyPrinter struct{}

func (lp *LegacyPrinter) PrintLegacy(message string) (string, error) {
	message = fmt.Sprintf("Legacy Printer: %s", message)
	fmt.Println(message)
	return message, nil
}

// ADAPTER
type PrinterAdapter struct {
	LegacyPrinter *LegacyPrinter
}

func (pa *PrinterAdapter) Print(message string) (string, error) {
	message, err := pa.LegacyPrinter.PrintLegacy(message)
	return message, err
}

// SERVIÇO QUE UTILIZA A INTERFACE LEGADA ATRAVÉS DO ADAPTER
type PrinterService struct {
	Printer IPrinter
}

func NewPrinterService(printer IPrinter) *PrinterService {
	return &PrinterService{Printer: printer}
}

func (ps *PrinterService) Print(message string) (string, error) {
	message, err := ps.Printer.Print(message)
	return message, err
}
