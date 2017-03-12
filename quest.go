package interact

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Question params
type Quest struct {
	Choices
	Default           Default
	Prefix            Prefix
	parent            *Question
	Options, Msg, Tag string
	Err               interface{}
	Resolve           BoolFunc
}

// Default answer value
type Default struct {
	Value   interface{}
	Text    interface{}
	Preview bool
}

// Question entity
type Question struct {
	Quest
	err           error
	reload        bool
	choices       bool
	response      string
	value         interface{}
	interact      *Interact
	parent        *Question
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

func (q *Question) ask() (err error) {
	context := &context{i: q.interact, q: q}
	if !context.i.skip {
		if err := context.method(q.Before); err != nil {
			return err
		}
		if !context.i.skip {
			if q.Prefix.Text != nil {
				q.print(q.Prefix.Text, " ")
			} else if q.parent != nil && q.parent.Prefix.Text != nil {
				q.print(q.parent.Prefix.Text, " ")
			} else if q.interact.Prefix.Text != nil {
				fmt.Print(q.interact.Prefix.Text, " ")
			}
			if q.Msg != "" {
				q.print(q.Msg, " ")
			}
			if q.Options != "" {
				q.print(q.Options, " ")
			}
			if q.Default.Preview && q.Default.Value != nil && q.Default.Text != nil {
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
					if q.Resolve(context) {
						for _, s := range q.Subs {
							s.interact = q.interact
							s.parent = q
							s.ask()
						}
					}
				} else {
					for _, s := range q.Subs {
						s.interact = q.interact
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
			if q.reload {
				q.reload = false
				return q.ask()
			}
			if err := context.method(q.After); err != nil {
				return err
			}
		} else {
			context.i.skip = false
		}
	} else {
		context.i.skip = false
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
	if q.Prefix.Writer != nil {
		fmt.Fprint(q.Prefix.Writer, v...)
	} else if q.parent != nil && q.parent.Prefix.Writer != nil {
		fmt.Fprint(q.parent.Prefix.Writer, v...)
	} else if q.interact != nil && q.interact.Prefix.Writer != nil {
		fmt.Fprint(q.interact.Prefix.Writer, v...)
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
	if q.Err != nil {
		q.print(q.Err, " ")
	} else if q.interact.Err != nil {
		q.print(q.interact.Err, " ")
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
