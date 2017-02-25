package interact

import (
	"time"
	"strconv"
	"errors"
)

type (
	Value interface {
		Int() int64
		Bool() bool
		String() string
		Float() float64
		Time() time.Duration
		Raw() interface{}
	}
	value struct {
		answer string
		choice interface{}
		err error
	}
)

func (v *value) Int() (value int64) {
	if v.choice != 0 {
		switch v.choice.(type) {
		case string:
			value, _ = strconv.ParseInt(v.choice.(string), 10, 64);
		case float64:
			value = int64(v.choice.(float64))
		case int:
			value = int64(v.choice.(int))
		default:
			v.err = errors.New("conversion as int failed")
		}
	}else {
		value, _ = strconv.ParseInt(v.answer, 10, 64)
	}
	return
}

func (v *value) Float() (value float64) {
	if v.choice != nil{
		switch v.choice.(type) {
		case string:
			value, _ = strconv.ParseFloat(v.choice.(string), 64);
		case float64:
			value = v.choice.(float64)
		case int:
			value = float64(v.choice.(int))
		default:
			v.err = errors.New("conversion as uint failed")
		}
	}else {
		value, _ = strconv.ParseFloat(v.answer, 64)
	}
	return
}

func (v *value) Time() time.Duration {
	if v.choice != nil{
		var cast int64
		switch v.choice.(type) {
		case string:
			cast, _ = strconv.ParseInt(v.choice.(string), 10, 64);
		case float64:
			cast = int64(v.choice.(float64))
		case int:
			cast = int64(v.choice.(int))
		default:
			v.err = errors.New("conversion as time duration failed")
		}
		return time.Duration(int64(cast))
	}
	if value, err := strconv.ParseUint(v.answer, 10, 64); err == nil {
		return time.Duration(value)
	}
	return time.Duration(0)
}

func (v *value) Bool() (value bool) {
	if v.choice != nil{
		switch v.choice.(type) {
		case bool:
			value, _ = strconv.ParseBool(v.answer)
		default:
			v.err = errors.New("conversion as bool failed")
		}
		return
	}
	if v.answer == "y" || v.answer == "yes" {
		return true
	} else if v.answer == "n" || v.answer == "no" {
		return false
	}
	value, _ = strconv.ParseBool(v.answer)
	return
}

func (v *value) String() (value string) {
	if v.choice != nil{
		switch v.choice.(type) {
		case string:
			value, _ = v.choice.(string);
		case int:
			value = strconv.Itoa(v.choice.(int))
		case float64:
			value = strconv.FormatFloat(v.choice.(float64), 'f', 2, 64)
		case bool:
			value = strconv.FormatBool(v.choice.(bool))
		default:
			v.err = errors.New("conversion as string failed")
		}
		return
	}
	return v.answer
}

func (v *value) Raw() interface{}{
	if v.choice != nil{
		return v.choice
	}
	return v.answer
}