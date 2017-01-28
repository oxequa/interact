package interact

import (
	"io"
)

// Interact interface
type I interface {
	ask(*Context) error
	quest() *Interact
	context() *Context
}

// Interact element
type Interact struct {
	prefix
	Questions     []*Question
	After, Before ErrorFunc
}

// Questions prefix
type prefix struct {
	Writer io.Writer
	Text   interface{}
}

// Run a questions list
func Run(i I) (*Interact, error) {
	context := i.context()
	if err := i.ask(context); err != nil {
		return nil, err
	}
	return i.quest(), nil
}

func (i *Interact) context() *Context {
	return &Context{interact: i}
}

func (i *Interact) quest() *Interact {
	return i
}

func (i *Interact) ask(c *Context) (err error) {
	for _, q := range i.Questions {
		q.parent = i
		c.quest = q
		if err = q.ask(c); err != nil {
			return err
		}
	}
	return nil
}
