package fu

import "reflect"

/*
Option returns first option of the required type
*/
func Option(t interface{}, o interface{}) reflect.Value {
	xs := reflect.ValueOf(o)
	tv := reflect.ValueOf(t)
	for i := 0; i < xs.Len(); i++ {
		x := xs.Index(i)
		if x.Kind() == reflect.Interface {
			x = x.Elem()
		}
		if x.Type() == tv.Type() {
			return x
		}
	}
	return tv
}

func Select[T comparable](selector T, opts []interface{}) T {
	for _, a := range opts {
		if v, ok := a.(T); ok {
			return v
		}
	}
	return selector
}
