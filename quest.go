package interact

import (
	"fmt"
	"strings"
	"strconv"
	"time"
	"bufio"
	"os"
)

type Quest struct {
	parent 	  *Interact
	resp      string
	SubQuest  []*Quest

	*Q
	Filter    func() 	// quests conditions
	Validate  func() 	// validate output
}

type Q struct {
	Text,Err,Response interface{}
	Default D
}

type D struct {
	Text, Value interface{}
}

func (q *Quest) quest() *Interact{
	i := Interact{}
	i.Questions = append(i.Questions, q)
	return &i
}

func (q *Quest) ask() (err error){
	if q.parent != nil && q.parent.W != nil{
		fmt.Fprint(q.parent.W,q.parent.T," ")
		fmt.Fprint(q.parent.W,q.Text,": ")
	}else {
		q.print(q.Text, ": ")
	}
	if q.Default.Value != nil{
		fmt.Print(q.Default.Text," ")
	}
	if err = q.wait(); err != nil {
		return err
	}
	if err = q.response(); err != nil {
		if q.Err != nil {
			fmt.Print(q.Err," ")
		}
		return q.ask()
	}
	return nil
}

func (q *Quest) wait() error {
	reader := bufio.NewReader(os.Stdin)
	r, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	q.resp = r[:len(r)-1]
	return nil
}

func (q *Quest) response() error {
	var v interface{}
	var err error

	switch q.Response.(type) {
	case uint:
		if v, err = strconv.ParseUint(q.resp, 10, 32); err == nil{
			q.Response = uint(v.(uint64))
		}
	case uint8:
		if v, err = strconv.ParseUint(q.resp, 10, 8); err == nil{
			q.Response = uint8(v.(uint64))
		}
	case uint16:
		if v, err = strconv.ParseUint(q.resp, 10, 16); err == nil{
			q.Response = uint16(v.(uint64))
		}
	case uint32:
		if v, err = strconv.ParseUint(q.resp, 10, 32); err == nil{
			q.Response = uint32(v.(uint64))
		}
	case uint64:
		q.Response, err = strconv.ParseUint(q.resp, 10, 64)
	case int:
		if v, err = strconv.ParseInt(q.resp, 10, 32); err == nil{
			q.Response = int(v.(int64))
		}
	case int8:
		if v, err = strconv.ParseInt(q.resp, 10, 8); err == nil{
			q.Response = int8(v.(int64))
		}
	case int16:
		if v, err = strconv.ParseInt(q.resp, 10, 16); err == nil{
			q.Response = int16(v.(int64))
		}
	case int32:
		if v, err = strconv.ParseInt(q.resp, 10, 32); err == nil{
			q.Response = int32(v.(int64))
		}
	case int64:
		q.Response, err = strconv.ParseInt(q.resp, 10, 64)
	case float32:
		if v, err = strconv.ParseFloat(q.resp, 64); err == nil{
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
		if v, err = strconv.ParseUint(q.resp, 10, 64); err == nil{
			q.Response = time.Duration(v.(uint64)) * time.Second
		}
	case string:
	default:
		q.Response = strings.ToLower(strings.TrimSpace(q.resp))
	}
	return err
}

func (q *Quest)print(a ...interface{}){
	if q.parent != nil && q.parent.W != nil {
		fmt.Fprint(q.parent.W, a...)
	}else if q.parent != nil && q.parent.W == nil {
		fmt.Print(q.parent.T)
		fmt.Print(a...)
	}else{
		fmt.Print(a...)
	}

}