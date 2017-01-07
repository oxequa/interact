package interact

import (
	"io"
)

type I interface {
	ask() error
	quest() *Interact
}

type Interact struct {
	Prefix
	Errors, Name      interface{}
	Questions []*Quest
}

type Context struct {
	error
	interact *Interact
	Quest *Quest
}

type Prefix struct {
	W io.Writer
	T interface{}
}

func Start(i I) (*Interact,error){
	if err := i.ask(); err != nil{
		return nil,err
	}
	return i.quest(),nil
}

func (i *Interact) quest() *Interact{
	return i
}

func (i *Interact) ask() (err error){
	for _, q := range i.Questions {
		q.parent = i
		if err = q.ask(); err != nil{
			return err
		}
	}
	return nil
}