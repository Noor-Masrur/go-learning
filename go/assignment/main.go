package main

import "fmt"

type square struct {
	side float64
}

type triangle struct {
	base float64
}

type shape interface {
	area() float64
}

func main() {
	s := square{side: 10}
	t := triangle{base: 10}
	fmt.Println(s.area())
	fmt.Println(t.area())
}

func (s square) area() float64 {
	return s.side * s.side
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.base
}
