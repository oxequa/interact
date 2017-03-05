package interact

import (
	"errors"
	"strconv"
	"time"
)

type (
	Cast interface {
		Int() int64
		Bool() bool
		String() string
		Float() float64
		Time() time.Duration
		Raw() interface{}
	}
	cast struct {
		answer string
		value  interface{}
		err    error
	}
)

type (
	Qns interface {
		Size() int
		Anwer() []Cast
		Get(int) Context
		GetTag(string) Context
		List() []Context
		ListTag(string) []Context
	}
	qns struct {
		list []Context
	}
)

// Cast the answer as an int
func (c *cast) Int() (value int64) {
	if c.value != nil {
		switch c.value.(type) {
		case string:
			value, _ = strconv.ParseInt(c.value.(string), 10, 64)
		case float64:
			value = int64(c.value.(float64))
		case int:
			value = int64(c.value.(int))
		default:
			c.err = errors.New("conversion as int failed")
		}
	} else {
		value, _ = strconv.ParseInt(c.answer, 10, 64)
	}
	return value
}

// Cast the answer as a float
func (c *cast) Float() (value float64) {
	if c.value != nil {
		switch c.value.(type) {
		case string:
			value, _ = strconv.ParseFloat(c.value.(string), 64)
		case float64:
			value = c.value.(float64)
		case int:
			value = float64(c.value.(int))
		default:
			c.err = errors.New("conversion as uint failed")
		}
	} else {
		value, _ = strconv.ParseFloat(c.answer, 64)
	}
	return
}

// Cast the answer as a time duration
func (c *cast) Time() time.Duration {
	if c.value != nil {
		var cast int64
		switch c.value.(type) {
		case string:
			cast, _ = strconv.ParseInt(c.value.(string), 10, 64)
		case float64:
			cast = int64(c.value.(float64))
		case int:
			cast = int64(c.value.(int))
		default:
			c.err = errors.New("conversion as time duration failed")
		}
		return time.Duration(int64(cast))
	}
	if value, err := strconv.ParseUint(c.answer, 10, 64); err == nil {
		return time.Duration(value)
	}
	return time.Duration(0)
}

// Cast the answer as a bool
func (c *cast) Bool() (value bool) {
	if c.value != nil {
		//fmt.Println(c.value)
		switch c.value.(type) {
		case bool:
			value = c.value.(bool)
		default:
			c.err = errors.New("conversion as bool failed")
		}
		return
	}
	if c.answer == "y" || c.answer == "yes" {
		return true
	} else if c.answer == "n" || c.answer == "no" {
		return false
	}
	value, _ = strconv.ParseBool(c.answer)
	return
}

// Cast the answer as a string
func (c *cast) String() (value string) {
	if c.value != nil {
		switch c.value.(type) {
		case string:
			value, _ = c.value.(string)
		case int:
			value = strconv.Itoa(c.value.(int))
		case float64:
			value = strconv.FormatFloat(c.value.(float64), 'f', 2, 64)
		case bool:
			value = strconv.FormatBool(c.value.(bool))
		default:
			c.err = errors.New("conversion as string failed")
		}
		return
	}
	return c.answer
}

// Raw return the answer as an interface
func (c *cast) Raw() interface{} {
	if c.value != nil {
		return c.value
	}
	return c.answer
}

// Size return the questions number
func (q *qns) Size() int {
	return len(q.list)
}

// Anwer return the answers for each question
func (q *qns) Anwer() []Cast {
	anwer := []Cast{}
	for _, question := range q.list {
		anwer = append(anwer, question.Ans())
	}
	return anwer
}

// Get a single question by index
func (q *qns) Get(i int) Context {
	for index, question := range q.list {
		if index == i {
			return question
		}
	}
	return &context{}
}

// Get a single question by tag
func (q *qns) GetTag(t string) Context {
	for _, question := range q.list {
		if question.Tag() == t {
			return question
		}
	}
	return nil
}

// Get the questions list
func (q *qns) List() []Context {
	return q.list
}

// Get the quesion list by tag
func (q *qns) ListTag(t string) []Context {
	list := []Context{}
	for _, question := range q.list {
		if question.Tag() == t {
			list = append(list, question)
		}
	}
	return list
}
