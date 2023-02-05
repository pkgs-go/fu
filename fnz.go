package fu

/*
Fnz returns first non-empty/zero value
*/
func Fnz[T comparable](a ...T) T {
	var zero T
	for _, x := range a {
		if x != zero {
			return x
		}
	}
	return zero
}
