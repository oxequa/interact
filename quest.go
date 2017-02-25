package interact

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"errors"
)

// Question params
type Quest struct {
	Choices
	Default
	parent             *Question
	Options, Err, Msg string
	Resolve BoolFunc
}

// Question entity
type Question struct {
	Quest
	prefix
	err                    error
	choices                bool
	response               string
	choice                 interface{}
	parent                 *Interact
	Action                 InterfaceFunc
	Subs []*Question
	After, Before ErrorFunc
}

// Default options
type Default struct {
	Text   interface{}
	Status bool
}

// Choice option
type Choice struct {
	Text     string
	Response interface{}
}

// Choices list and prefix color
type Choices struct {
	Alternatives []Choice
	Color        func(...interface{}) string
}

func (q *Question) answer() interface{}{
	return value{answer:q.response, choice: q.choice, err: q.err}
}

func (q *Question) append(p prefix) {
	q.prefix = p
}

func (q *Question) father() model {
	return q.parent
}

func (q *Question) ask() (err error) {
	context := &context{model: q}
	if err := context.method(q.Before); err != nil{
		return err
	}
	if q.prefix.Text != nil{
		q.print(q.prefix.Text, " ")
	}else if q.parent != nil && q.parent.Text != nil {
		q.print(q.parent.Text, " ")
	}
	if q.Msg != "" {
		q.print(q.Msg, " ")
	}
	if q.Options != "" {
		q.print(q.Options, " ")
	}
	if q.Default.Status != false {
		q.print(q.Default.Text, " ")
	}
	if q.Alternatives != nil && len(q.Alternatives) > 0 {
		q.multiple()
	}
	if err = q.wait(); err != nil {
		return q.loop(err)
	}
	if q.Subs != nil && len(q.Subs) > 0 {
		if q.Resolve != nil {
			if q.Resolve(context){
				for _, s := range q.Subs {
					s.parent = q.parent
					s.ask()
				}
			}
		}else{
			for _, s := range q.Subs {
				s.parent = q.parent
				s.ask()
			}
		}
	}
	if q.Action != nil {
		if err := q.Action(context); err != nil {
			q.print(err, " ")
			return q.ask()
		}
	}
	if err := context.method(q.After); err != nil{
		return err
	}
	return nil
}

func (q *Question) wait() error {
	reader := bufio.NewReader(os.Stdin)
	if q.choices {
		q.print(q.color("?"), " ", "Answer", " ")
	}
	r, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	q.response = r[:len(r)-1]

	if len(q.response) == 0 && q.Default.Status {
		return nil
	}
	// multiple choice
	if q.choices {
		choice, err := strconv.ParseInt(q.response, 10, 64)
		if err != nil || int(choice) > len(q.Alternatives) {
			return errors.New("out of range")
		}
		q.choice = q.Alternatives[choice-1].Response
	}
	return nil
}

func (q *Question) print(v ...interface{}) {
	if q.parent != nil && q.parent.Writer != nil {
		fmt.Fprint(q.parent.Writer, v...)
	} else {
		fmt.Print(v...)
	}

}

func (q *Question) color(v ...interface{}) string{
	if q.Color != nil{
		return q.Color(v...)
	}
	return fmt.Sprint(v...)
}

func (q *Question) loop(err error) error {
	if q.Err != "" {
		q.print(q.Err, " ")
	}
	return q.ask()
}

func (q *Question) multiple() error{
	for index,i := range q.Alternatives {
		q.print("\n\t", q.color(index + 1, ") "), i.Text, " ")
	}
	q.choices = true
	q.print("\n")
	return nil
}