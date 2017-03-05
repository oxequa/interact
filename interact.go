package interact

import (
	"io"
)

type (
	Interview interface {
		//Next()
		//Prev()
		//Goto(interface{})
	}
	interview struct {
		index    int
		interact *Interact
		current  current
	}
	current struct {
		list  []*Question
		index int
	}
)

// Interact element
type Interact struct {
	skip          bool
	Prefix        Prefix
	Default       Default
	Questions     []*Question
	Err           interface{}
	After, Before ErrorFunc
}

// Questions prefix
type Prefix struct {
	Writer io.Writer
	Text   interface{}
}

// Run a questions list
func Run(i *Interact) error {
	if err := i.ask(); err != nil {
		return err
	}
	return nil
}

func New(i *Interact) Interview {
	return &interview{interact: i, current: current{index: 0, list: i.Questions}}
}

func (i *Interact) ask() (err error) {
	context := &context{i: i}
	if err := context.method(i.Before); err != nil {
		return err
	}
	if !context.i.skip {
		for index := range i.Questions {
			i.Questions[index].interact = i
			if err = i.Questions[index].ask(); err != nil {
				return err
			}
		}
		if err := context.method(i.After); err != nil {
			return err
		}
	}
	return nil
}
