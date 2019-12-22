package main

import (
	"fmt"

	cpu "./utils"
)

const (
	A = 5
	B = 6
)

type Shape interface {
	area() float64
}
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	cpu1 := cpu.Cpu{SP: 0x1ff}
	cpu2 := &cpu1
	fmt.Print()
	cpu2.Debug()

}
