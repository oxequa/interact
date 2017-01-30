package interact

import (
	"io"
)

type interact interface {
	ask() error
	context() Context
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
func Run(i interact) error {
	if err := i.ask(); err != nil {
		return err
	}
	return nil
}

func (i *Interact) append(p prefix) {
	i.prefix = p
}

func (i *Interact) father() model {
	return i
}

func (i *Interact) ask() (err error) {
	context := i.context()
	if err := i.Before(context); err != nil {
		return err
	}
	for index := range i.Questions {
		i.Questions[index].parent = i
		if err = i.Questions[index].ask(); err != nil {
			return err
		}
	}
	if err := i.After(context); err != nil {
		return err
	}
	return nil
}

func (i *Interact) context() Context {
	c := context{model: i}
	return &c
}

func (i *Interact) answer() interface{} {
	answers := []interface{}{}
	for _, q := range i.Questions {
		answers = append(answers, q.Quest.Response)
	}
	return answers
}
