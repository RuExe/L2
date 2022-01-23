package strategy

import "fmt"

type Strategy interface {
	execute(a, b int) int
}

type AddStrategy struct {
}

func (s AddStrategy) execute(a, b int) int {
	return a + b
}

type SubtractStrategy struct {
}

func (s SubtractStrategy) execute(a, b int) int {
	return a - b
}

type MultiplyStrategy struct {
}

func (s MultiplyStrategy) execute(a, b int) int {
	return a * b
}

type Context struct {
	strategy Strategy
}

func (c Context) executeStrategy(a, b int) int {
	return c.strategy.execute(a, b)
}

func main() {
	a, b := 1, 2

	s := 1
	var strat Strategy = AddStrategy{}
	switch s {
	case 2:
		strat = SubtractStrategy{}
	case 3:
		strat = MultiplyStrategy{}
	}

	context := Context{strat}
	fmt.Println(context.executeStrategy(a, b))
}
