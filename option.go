package fu

/*
Select returns the first option of the required type
*/
func Select[T any](selector T, opts []interface{}) T {
	for _, a := range opts {
		if v, ok := a.(T); ok {
			return v
		}
	}
	return selector
}
