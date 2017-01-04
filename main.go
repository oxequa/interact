package interact

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type interactive interface {
	Ask()
	Start()
}

type interact struct {
	interactive
	name      string
	questions []quest
}

type quest struct {
	parent    *interact
	field     string
	text      string
	expected  string
	validate  validate
	subQuest  []quest
	filter    func() // for subQuest
	inputType interface{}
	obj       interface{}
	def       interface{}
}

type msg struct {
	Name     string
	Field    string
	Text     string
	Expected string
	Def      interface{}
}

type validate struct {
	valid func()
	err   string
}

func New(n ...string) (i *interact){
	if(len(n)>0) {
		i = &interact{name:n[0]}
	}
	return i
}

func (i quest) Start() {

}

func (q quest) Ask() interface{} {
	q.quest()
	val, err := q.wait()
	if err != nil {
		return err
	}
	return val
}

func (q quest) quest() {
	v := reflect.ValueOf(msg{Name: q.parent.name, Field: q.field, Text: q.text, Expected: q.expected, Def: q.def})
	for i := 0; i < v.NumField()-1; i++ {
		f := v.Field(i)
		// check default value
		if f.CanInterface() && !empty(f.Interface()) {
			switch i {
			case 2:
				fmt.Print(" ", f, " ")
			case 4:
				fmt.Print("(", f, ")")
			case 0:
				fmt.Print("[", strings.ToUpper(f.String()), "]")
			default:
				fmt.Print("[", f.String(), "]")
			}
		}
	}
	def := v.Field(v.NumField() - 1)
	if def.Interface() != nil {
		fmt.Print(":", " ", "(", def, ")", " ")
	}
}

func (q quest) wait() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	val, err := reader.ReadString('\n')
	if err != nil {
		return val, err
	}
	return val[:len(val)-1], nil
}

func empty(x interface{}) bool {
	return x == nil || x == reflect.Zero(reflect.TypeOf(x)).Interface()
}