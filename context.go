package interact

type Context struct {
	interact *Interact
	quest    *Quest
}

type Validate func(*Context) interface{}

type Filter func(*Context) error
