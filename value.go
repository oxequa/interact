package interact

import (
	"errors"
	"strconv"
	"time"
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
		value  interface{}
		err    error
	}
)

func (v *value) Int() (value int64) {
	if v.value != nil {
		switch v.value.(type) {
		case string:
			value, _ = strconv.ParseInt(v.value.(string), 10, 64)
		case float64:
			value = int64(v.value.(float64))
		case int:
			value = int64(v.value.(int))
		default:
			v.err = errors.New("conversion as int failed")
		}
	} else {
		value, _ = strconv.ParseInt(v.answer, 10, 64)
	}
	return value
}

func (v *value) Float() (value float64) {
	if v.value != nil {
		switch v.value.(type) {
		case string:
			value, _ = strconv.ParseFloat(v.value.(string), 64)
		case float64:
			value = v.value.(float64)
		case int:
			value = float64(v.value.(int))
		default:
			v.err = errors.New("conversion as uint failed")
		}
	} else {
		value, _ = strconv.ParseFloat(v.answer, 64)
	}
	return
}

func (v *value) Time() time.Duration {
	if v.value != nil {
		var cast int64
		switch v.value.(type) {
		case string:
			cast, _ = strconv.ParseInt(v.value.(string), 10, 64)
		case float64:
			cast = int64(v.value.(float64))
		case int:
			cast = int64(v.value.(int))
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
	if v.value != nil {
		//fmt.Println(v.value)
		switch v.value.(type) {
		case bool:
			value = v.value.(bool)
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
	if v.value != nil {
		switch v.value.(type) {
		case string:
			value, _ = v.value.(string)
		case int:
			value = strconv.Itoa(v.value.(int))
		case float64:
			value = strconv.FormatFloat(v.value.(float64), 'f', 2, 64)
		case bool:
			value = strconv.FormatBool(v.value.(bool))
		default:
			v.err = errors.New("conversion as string failed")
		}
		return
	}
	return v.answer
}

func (v *value) Raw() interface{} {
	if v.value != nil {
		return v.value
	}
	return v.answer
}
