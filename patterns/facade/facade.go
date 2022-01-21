package facade

import "fmt"

type (
	Facade struct {
		subSystem1 *SubSystem1
		subSystem2 *SubSystem2
	}

	SubSystem1 struct {
	}

	SubSystem2 struct {
	}
)

func NewFacade(s1 *SubSystem1, s2 *SubSystem2) *Facade {
	return &Facade{
		subSystem1: s1,
		subSystem2: s2,
	}
}

func (f *Facade) Operation() string {
	res := "Facade initializes subsystems:\n"
	res += fmt.Sprintln(f.subSystem1.Operation())
	res += fmt.Sprintln(f.subSystem2.Operation())
	return res
}

func (s *SubSystem1) Operation() string {
	return "Subsystem1: Ready!"
}

func (s *SubSystem2) Operation() string {
	return "Subsystem2: Get ready!"
}
