package derived

import "example.com/crtp/base"

// DerivedA is an example implementation
type DerivedA struct {
	base.Base[DerivedA]
}

func (DerivedA) Name() string {
	return "Reasonable person"
}

func (DerivedA) Message() string {
	return "Please don't do this in real code"
}

// DerivedB is an example implementation
type DerivedB struct {
	base.Base[DerivedB]
}

func (DerivedB) Name() string {
	return "Edu"
}

func (DerivedB) Message() string {
	return "I'm going to do this everywhere!"
}
