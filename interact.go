package interact

import (
	"io"
)

type I interface {
	ask(*Context) error
	quest() *Interact
	context() *Context
}

type Interact struct {
	Prefix
	Errors, Name      interface{}
	Questions []*Quest
}

type Prefix struct {
	W io.Writer
	T interface{}
}

func Run(i I) (*Interact,error){
	context := i.context()
	if err := i.ask(context); err != nil{
		return nil,err
	}
	return i.quest(),nil
}

func(i *Interact) context() *Context{
	return &Context{interact:i}
}

func (i *Interact) quest() *Interact{
	return i
}

func (i *Interact) ask(c *Context) (err error){
	for _, q := range i.Questions {
		q.parent = i
		c.quest = q
		if err = q.ask(c); err != nil{
			return err
		}
	}
	return nil
}
