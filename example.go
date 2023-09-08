package main

import (
	"fmt"
)

// ------------ Define the embedded type ------------
type derived interface {
	Message() string
}

type Base[T derived] struct{}

func (b Base[T]) Speak() {
	var t T
	fmt.Printf(t.Message())
}

// ------------ Define the implementation ------------

type Derived struct {
	Base[Derived]
}

func (Derived) Message() string {
	return "Please don't do this in real code"
}

// ------------------- Example usage -------------------

func main() {
	var d Derived

	d.Speak() // -> Prints 'Please don't do this in real code'
}
