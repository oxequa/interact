package interact

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	resp    string
	parent  *Interact
	choices bool
	*Quest
	Subs
	Action                 InterfaceFunc
	After, Before, Resolve ErrorFunc
}

type Choices struct {
	Alternatives []Choice
	Prefix string
	Color    func(...interface{}) string
}

type Choice struct {
	Response interface{}
	Text     string
}

type Quest struct {
	Response interface{}
	Default
	Options, Err, Text string
	Choices
}

type Default struct {
	Text   interface{}
	Status bool
}

// Related sub questions
type Subs struct {
	Questions []*Question
	Resolve   ErrorFunc // quests conditions for sub questions
}

func (q *Question) context() *Context {
	i := Interact{}
	return &Context{interact: &i, quest: q}
}

func (q *Question) quest() *Interact {
	i := Interact{}
	i.Questions = append(i.Questions, q)
	return &i
}

func (q *Question) ask(c *Context) (err error) {
	if q.parent != nil && q.parent.Text != nil {
		q.print(q.parent.Text, " ")
	}
	if q.Text != "" {
		q.print(q.Text, " ")
	}
	if q.Options != "" {
		q.print(q.Options, " ")
	}
	if q.Default.Status != false {
		q.print(q.Default.Text, " ")
	}
	if q.Alternatives != nil && len(q.Alternatives) > 0 {
		for index, i := range q.Alternatives {
			q.print("\n", q.Prefix, q.Color(index+1, ") "), i.Text, " ")
		}
		q.choices = true
		q.print("\n")
	}
	if err = q.wait(); err != nil {
		return q.loop(err, c)
	}
	if err = q.response(); err != nil {
		return q.loop(err, c)
	}
	if err := q.Action(c); err != nil {
		q.print(err, " ")
		return q.ask(c)
	}
	if q.Subs.Questions != nil && len(q.Subs.Questions) > 0 {
		if err := q.Subs.Resolve(c); err != nil {
			for _, s := range q.Subs.Questions {
				s.ask(c)
			}
		}

	}
	return nil
}

func (q *Question) wait() error {
	reader := bufio.NewReader(os.Stdin)
	if q.choices {
		q.print(q.Color("?")," ","Answer"," ")
	}
	r, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	q.resp = r[:len(r)-1]
	return nil
}

func (q *Question) response() error {
	var v interface{}
	var err error

	// dafault response
	if len(q.resp) == 0 && q.Default.Status{
		return nil
	}
	// multiple choice
	if q.choices{
		q.Response, err = strconv.ParseInt(q.resp, 10, 64)
		if (err != nil){
			return err
		}
		q.Response = q.Alternatives[q.Response.(int64)-1].Response
	}

	switch q.Response.(type) {
	case uint:
		if v, err = strconv.ParseUint(q.resp, 10, 32); err == nil {
			q.Response = uint(v.(uint64))
		}
	case uint8:
		if v, err = strconv.ParseUint(q.resp, 10, 8); err == nil {
			q.Response = uint8(v.(uint64))
		}
	case uint16:
		if v, err = strconv.ParseUint(q.resp, 10, 16); err == nil {
			q.Response = uint16(v.(uint64))
		}
	case uint32:
		if v, err = strconv.ParseUint(q.resp, 10, 32); err == nil {
			q.Response = uint32(v.(uint64))
		}
	case uint64:
		q.Response, err = strconv.ParseUint(q.resp, 10, 64)
	case int:
		if v, err = strconv.ParseInt(q.resp, 10, 32); err == nil {
			q.Response = int(v.(int64))
		}
	case int8:
		if v, err = strconv.ParseInt(q.resp, 10, 8); err == nil {
			q.Response = int8(v.(int64))
		}
	case int16:
		if v, err = strconv.ParseInt(q.resp, 10, 16); err == nil {
			q.Response = int16(v.(int64))
		}
	case int32:
		if v, err = strconv.ParseInt(q.resp, 10, 32); err == nil {
			q.Response = int32(v.(int64))
		}
	case int64:
		q.Response, err = strconv.ParseInt(q.resp, 10, 64)
	case float32:
		if v, err = strconv.ParseFloat(q.resp, 64); err == nil {
			q.Response = float32(v.(float64))
		}
	case float64:
		q.Response, err = strconv.ParseFloat(q.resp, 64)
	case bool:
		if q.resp == "y" || q.resp == "yes" {
			q.Response = true
		} else if q.resp == "n" || q.resp == "no" {
			q.Response = false
		} else {
			q.Response, err = strconv.ParseBool(q.resp)
		}
	case time.Duration:
		if v, err = strconv.ParseUint(q.resp, 10, 64); err == nil {
			q.Response = time.Duration(v.(uint64)) * time.Second
		}
	case string:
	default:
		q.Response = strings.ToLower(strings.TrimSpace(q.resp))
	}
	return err
}

func (q *Question) print(a ...interface{}) {
	if q.parent != nil && q.parent.Writer != nil {
		fmt.Fprint(q.parent.Writer, a...)
	} else {
		fmt.Print(a...)
	}

}

func (q *Question) loop (err error, c *Context) error{
	if q.Err != "" {
		q.print(q.Err, " ")
	}
	return q.ask(c)
}