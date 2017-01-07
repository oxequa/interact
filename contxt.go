package interact

type Context struct {
	interact *Interact
	quest *Quest
}

type ContextFunc func(*Context) interface{}
