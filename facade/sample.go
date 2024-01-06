package facade

import "fmt"

// Subsistema 1
type System1 struct {
}

func (s *System1) Execute(param string) (string, error) {
	result := fmt.Sprintf("Subsistema 1: Operação com parâmetro %s", param)
	return result, nil
}

// Subsistema 2
type System2 struct {
}

func (s *System2) Execute(param string) (string, error) {
	result := fmt.Sprintf("Subsistema 2: Operação com parâmetro %s", param)
	return result, nil
}

type Facade struct {
	system1 *System1
	system2 *System2
}

func NewFacade() *Facade {
	return &Facade{
		system1: &System1{},
		system2: &System2{},
	}
}

func (f *Facade) Execute() error {
	fmt.Println("Facade: Iniciando operação complexa")
	result, err := f.system1.Execute("param1")
	if err != nil {
		return err
	}
	fmt.Println(result)

	result, err = f.system2.Execute("param2")
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}
