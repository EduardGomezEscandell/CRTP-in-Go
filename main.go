package main

import "example.com/crtp/derived"

func main() {
	var (
		a derived.Sceptic
		b derived.Believer
	)

	a.Speak()
	b.Speak()
}
