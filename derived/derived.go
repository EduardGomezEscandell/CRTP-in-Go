package derived

import "example.com/crtp/base"

// Sceptic is an example implementation
type Sceptic struct {
	base.Person[Sceptic]
}

func (Sceptic) Name() string {
	return "Reasonable person"
}

func (Sceptic) Message() string {
	return "Please don't do this in real code"
}

// Believer is an example implementation
type Believer struct {
	base.Person[Believer]
}

func (Believer) Name() string {
	return "Intern"
}

func (Believer) Message() string {
	return "I'm going to do this everywhere!"
}
