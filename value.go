package interact

type (
	Value interface {
		Int() int
		Uint() uint
		Bool() bool
		String() string
		Raw() interface{}
	}
	value struct {
		val interface{}
	}
)

func (v *value) Int() int {
	value, _ := v.val.(int)
	return value
}

func (v *value) Uint() uint {
	value, _ := v.val.(uint)
	return value
}

func (v *value) Bool() bool {
	value, _ := v.val.(bool)
	return value
}

func (v *value) String() string {
	value, _ := v.val.(string)
	return value
}

func (v *value) Raw() interface{} {
	return v.val
}
