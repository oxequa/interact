package interact

import "io"

type Context struct {
	interact *Interact
	quest    *Question
}

type InterfaceFunc func(*Context) interface{}

type ErrorFunc func(*Context) error

func (c *Context) Prefix(w io.Writer, t interface{}) error {
	c.interact.prefix.Writer = w
	c.interact.prefix.Text = t
	return nil
}

func (c *Context) Response() interface{} {
	return c.quest.Response
}

func (c *Context) ResponseBool() bool {
	return c.quest.Response.(bool)
}

func (c *Context) ResponseString() string {
	return c.quest.Response.(string)
}

func (c *Context) ResponseInt() int {
	return c.quest.Response.(int)
}

func (c *Context) ResponseChoice() int {
	return c.quest.Response.(int)
}
