package interact

type I interface {
	ask() error
}

type Interact struct {
	Prefix      interface{}
	Questions []*Quest
}

type Context struct {
	error
	interact *Interact
	Quest *Quest
}

func Start(i I){
	i.ask()
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