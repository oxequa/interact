package interact

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Question params
type Quest struct {
	Choices
	Default           Default
	parent            *Question
	Options, Err, Msg string
	Resolve           BoolFunc
}

type Default struct {
	Value   interface{}
	Preview bool
}

// Question entity
type Question struct {
	Quest
	prefix
	err           error
	choices       bool
	response      string
	value         interface{}
	parent        model
	Action        InterfaceFunc
	Subs          []*Question
	After, Before ErrorFunc
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

func (q *Question) answer() interface{} {
	return value{answer: q.response, value: q.value, err: q.err}
}

func (q *Question) append(p prefix) {
	q.prefix = p
}

func (q *Question) father() model {
	return q.parent
}

func (q *Question) writer() io.Writer {
	return q.Writer
}

func (q *Question) lead() interface{} {
	return q.Text
}

func (q *Question) ask() (err error) {
	context := &context{model: q}
	if err := context.method(q.Before); err != nil {
		return err
	}
	if q.lead() != nil {
		q.print(q.lead(), " ")
	} else if q.parent != nil && q.parent.lead() != nil {
		q.print(q.parent.lead(), " ")
	}
	if q.Msg != "" {
		q.print(q.Msg, " ")
	}
	if q.Options != "" {
		q.print(q.Options, " ")
	}
	if q.Default.Preview && q.Default.Value != nil {
		q.print(q.Default.Value, " ")
	}
	if q.Alternatives != nil && len(q.Alternatives) > 0 {
		q.multiple()
	}
	if err = q.wait(); err != nil {
		return q.loop(err)
	}
	if q.Subs != nil && len(q.Subs) > 0 {
		if q.Resolve != nil {
			if q.Resolve(context) {
				for _, s := range q.Subs {
					s.parent = q
					s.ask()
				}
			}
		} else {
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
	if err := context.method(q.After); err != nil {
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

	if len(q.response) == 0 && q.Default.Value != nil {
		q.value = q.Default.Value
		return nil
	} else if len(q.response) == 0 {
		return errors.New("Answer invalid")
	}

	// multiple choice
	if q.choices {
		choice, err := strconv.ParseInt(q.response, 10, 64)
		if err != nil || int(choice) > len(q.Alternatives) {
			return errors.New("out of range")
		}
		q.value = q.Alternatives[choice-1].Response
	}
	return nil
}

func (q *Question) print(v ...interface{}) {
	if q.writer() != nil {
		fmt.Fprint(q.writer(), v...)
	} else if q.parent != nil && q.parent.writer() != nil {
		fmt.Fprint(q.parent.writer(), v...)
	} else {
		fmt.Print(v...)
	}

}

func (q *Question) color(v ...interface{}) string {
	if q.Color != nil {
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

func (q *Question) multiple() error {
	for index, i := range q.Alternatives {
		q.print("\n\t", q.color(index+1, ") "), i.Text, " ")
	}
	q.choices = true
	q.print("\n")
	return nil
}
