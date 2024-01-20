package polymophism

import "fmt"

type Shape interface {
	Rander()
}

// Implements Rander
type Circle struct{}

func (c Circle) Rander() {
	fmt.Println("Circle is getting rendered!")
}

// Implements Rander
type Square struct{}

func (s Square) Rander() {
	fmt.Println("Sqaure is getting rendered!")
}