package interact

import (
	"io"
)

type interact interface {
	ask() error
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
	context := &context{model: i}
	if err := context.method(i.Before); err != nil{
		return err
	}
	for index := range i.Questions {
		i.Questions[index].parent = i
		if err = i.Questions[index].ask(); err != nil {
			return err
		}
	}
	if err := context.method(i.After); err != nil{
		return err
	}
	return nil
}

func (i *Interact) answer() interface{} {
	answers := []value{}
	for _, q := range i.Questions {
		answers = append(answers, value{answer:q.response, choice: q.choice, err: q.err})
	}
	return answers
}

func (i *Interact) writer() io.Writer{
	return i.Writer
}

func (i *Interact) lead() interface{}{
	return i.Text
}