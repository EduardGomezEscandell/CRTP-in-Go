package base

import "fmt"

type derived interface {
	Name() string
	Message() string
}

type Base[T derived] struct{}

func (b Base[T]) Speak() {
	var t T
	fmt.Printf("%s says: %s\n", t.Name(), t.Message())
}
