package interact

import (
	"errors"
	"io"
)

type (
	Context interface {
		Skip()
		Reload()
		SetPrfx(io.Writer, interface{})
		SetDef(interface{}, interface{}, bool)
		SetErr(interface{})
		Ans() Cast
		Def() Cast
		Err() error
		Parent() Context
		Prfx() Cast
		Qns() Qns
		Tag() string
		Quest() string
	}
	context struct {
		i *Interact
		q *Question
	}
)

type ErrorFunc func(Context) error

type BoolFunc func(Context) bool

type InterfaceFunc func(Context) interface{}

func (c *context) Parent() Context {
	if c.q.parent != nil {
		return &context{i: c.i, q: c.q.parent}
	}
	return &context{i: c.i}
}

func (c *context) Ans() Cast {
	if c.q != nil {
		return &cast{answer: c.q.response, value: c.q.value, err: c.q.err}
	}
	return &cast{}
}

func (c *context) Skip() {
	c.i.skip = true
}

func (c *context) Reload() {
	if c.q != nil {
		c.q.reload = true
	}
}

func (c *context) Err() error {
	if c.q.Err != nil {
		return errors.New(c.q.Err.(string))
	} else if c.i.Err != nil {
		return errors.New(c.i.Err.(string))
	}
	return nil
}

func (c *context) Def() Cast {
	if c.q != nil {
		return &cast{value: c.q.Default.Value}
	}
	return &cast{value: c.i.Default.Value}
}

func (c *context) Prfx() Cast {
	if c.q != nil && c.q.Prefix.Text != nil {
		return &cast{value: c.q.Prefix.Text}
	}
	return &cast{value: c.i.Prefix.Text}
}

func (c *context) Qns() Qns {
	list := []Context{}
	if c.q != nil {
		for _, q := range c.q.Subs {
			list = append(list, &context{i: c.i, q: q})
		}
	} else {
		for _, q := range c.i.Questions {
			list = append(list, &context{i: c.i, q: q})
		}
	}
	return &qns{list: list}
}

func (c *context) Tag() string {
	if c.q != nil {
		return c.q.Tag
	}
	return ""
}

func (c *context) Quest() string {
	if c.q != nil {
		return c.q.Quest.Msg
	}
	return ""
}

func (c *context) SetPrfx(w io.Writer, t interface{}) {
	if c.q != nil {
		c.q.Prefix = Prefix{w, t}
		return
	}
	c.i.Prefix = Prefix{w, t}
	return
}

func (c *context) SetDef(v interface{}, t interface{}, p bool) {
	if c.q != nil {
		c.q.Default = Default{v, t, p}
		return
	}
	c.i.Default = Default{v, t, p}
	return
}

func (c *context) SetErr(e interface{}) {
	if c.q != nil {
		c.q.Err = e
		return
	}
	c.i.Err = e
	return
}

func (c *context) method(f ErrorFunc) error {
	if f != nil {
		if err := f(c); err != nil {
			return err
		}
	}
	return nil
}
