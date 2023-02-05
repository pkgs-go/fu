package fu

/*
FirstOr returns first value from the range or default one
*/
func FirstOr[T comparable](v T, a ...T) T {
	if len(a) > 0 {
		return a[0]
	}
	return v
}
