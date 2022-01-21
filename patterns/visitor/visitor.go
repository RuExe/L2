package visitor

import (
	"fmt"
	"math"
)

type Shape interface {
	accept(v Visitor)
}

type Circle struct {
	radius int
}

func (c *Circle) Accept(v Visitor) {
	v.visitCircle(*c)
}

type Square struct {
	side int
}

func (c *Square) Accept(v Visitor) {
	v.visitSquare(*c)
}

type Visitor interface {
	visitCircle(c Circle)
	visitSquare(r Square)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForCircle(s *Circle) {
	fmt.Printf("Calculating area for circle %v\n", math.Pi*float64(s.radius*s.radius))
}
func (a *areaCalculator) visitForrectangle(s *Square) {
	fmt.Printf("Calculating area for rectangle %v\n", s.side*s.side)
}
