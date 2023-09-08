package main

import "example.com/crtp/derived"

func main() {
	var (
		a derived.Derived
	)

	a.Speak() // -> Prints 'Edu says: Please don't do this in real code'
}
