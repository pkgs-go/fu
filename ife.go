package fu

/*
Ife returns value selected by logical expression
*/
func Ife[T any](expr bool, onTrue, onFalse T) T {
	if expr {
		return onTrue
	}
	return onFalse
}
