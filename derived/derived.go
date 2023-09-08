package derived

import "example.com/crtp/base"

// Derived is an example implementation
type Derived struct {
	base.Base[Derived]
}

func (Derived) Name() string {
	return "Edu"
}

func (Derived) Message() string {
	return "Please don't do this in real code"
}
