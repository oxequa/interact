package interact

import (
	"io"
)

type (
	Context interface {
		Parent() Context
		Answer() Value
		Answers() []Value
		Input() Value
		Prefix(io.Writer, interface{})
	}
	context struct {
		model
	}
)

type (
	model interface {
		father() model
		append(prefix)
		writer() io.Writer
		lead() interface{}
		answer() interface{}
	}
)

type ErrorFunc func(Context) error

type BoolFunc func(Context) bool

type InterfaceFunc func(Context) interface{}

func (c *context) Parent() Context {
	return &context{model: c.father()}
}

func (c *context) Answer() Value {
	answ, _ := c.answer().(value)
	return &answ
}

func (c *context) Answers() (v []Value) {
	answers, _ := c.answer().([]value)
	for index := range answers {
		v = append(v, &answers[index])
	}
	return v
}

func (c *context) Input() Value {
	answ := c.answer().(value)
	return &answ
}

func (c *context) Prefix(w io.Writer, t interface{}) {
	p := prefix{w, t}
	c.append(p)
}

func (c *context) method(f ErrorFunc) error {
	if f != nil {
		if err := f(c); err != nil {
			return err
		}
	}
	return nil
}
