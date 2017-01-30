package interact

import (
	"io"
)

type (
	Context interface {
		Parent() Context
		Int() int
		Uint() uint
		Bool() bool
		String() string
		Value() interface{}
		Prefix(io.Writer, interface{})
	}
	context struct {
		model
	}
)

type model interface {
	father() model
	append(prefix)
	context() Context
	answer() interface{}
}

type ErrorFunc func(Context) error

type InterfaceFunc func(Context) interface{}

func (c *context) Int() int {
	return c.answer().(int)
}

func (c *context) Uint() uint {
	return c.answer().(uint)
}

func (c *context) Bool() bool {
	return c.answer().(bool)
}

func (c *context) String() string {
	return c.answer().(string)
}

func (c *context) Parent() Context {
	clone := c
	clone.model = clone.father()
	return clone
}

func (c *context) Value() interface{} {
	return c.answer()
}

func (c *context) Prefix(w io.Writer, t interface{}) {
	p := prefix{w,t}
	c.append(p)
}
