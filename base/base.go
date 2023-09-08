package base

import "fmt"

type individual interface {
	Name() string
	Message() string
}

type Person[T individual] struct{}

func (b Person[T]) Speak() {
	var t T
	fmt.Printf("%s says: %s\n", t.Name(), t.Message())
}
