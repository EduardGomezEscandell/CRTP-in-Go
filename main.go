package main

import "example.com/crtp/derived"

func main() {
	var (
		a derived.DerivedA
		b derived.DerivedB
	)

	a.Speak()
	b.Speak()
}
