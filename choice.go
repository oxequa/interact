package interact

type Choice struct {
	prefix
	parent   *Question
	Text     string
	Response interface{}
}

type Choices struct {
	Alternatives []Choice
	Color        func(...interface{}) string
}

func (ch *Choice) context() Context {
	c := context{model: ch}
	return &c
}

func (ch *Choice) answer() interface{} {
	return ch.Response
}

func (ch *Choice) append(p prefix) {
	ch.prefix = p
}

func (ch *Choice) father() model {
	return ch.parent
}
