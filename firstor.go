package fu

/*
FirstOr returns first value from the range or default default one
*/
func FirstOr[T comparable](v T, a ...T) T {
	for _, x := range a {
		return x
	}
	return v
}
