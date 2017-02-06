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

type(
	model interface {
		father() model
		append(prefix)
		answer() interface{}
	}
	response struct {
		input interface{}
		answer interface{}
	}
)

type ErrorFunc func(Context) error

type InterfaceFunc func(Context) interface{}

func (c *context) Parent() Context {
	clone := c
	clone.model = clone.father()
	return clone
}

func (c *context) Answer() Value {
	answ := c.answer().(response)
	return &value{val: answ.answer}
}

func (c *context) Answers() (v []Value) {
	answers, _ := c.answer().([]response)
	for index := range answers{
		v = append(v,&value{val:answers[index].answer})
	}
	return v
}

func (c *context) Input() Value {
	answ := c.answer().(response)
	return &value{val: answ.input}
}

func (c *context) Prefix(w io.Writer, t interface{}) {
	p := prefix{w,t}
	c.append(p)
}
