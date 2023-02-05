package fu

type ConstContainer[T comparable] interface {
	Contains(val T) bool
}

type sliceOf[T comparable] ([]T)

func Slice[T comparable](t []T) ConstContainer[T] {
	return sliceOf[T](t)
}

func (self sliceOf[T]) Contains(val T) bool {
	for _, a := range self {
		if a == val {
			return true
		}
	}
	return false
}

type mapOf[K, V comparable] (map[K]V)

func MapOf[K, V comparable](m map[K]V) ConstContainer[V] {
	return mapOf[K, V](m)
}

func (self mapOf[any, T]) Contains(val T) bool {
	for _, a := range self {
		if a == val {
			return true
		}
	}
	return false
}

func (self mapOf[T, any]) Keys() []T {
	keys := make([]T, 0, len(self))
	for k := range self {
		keys = append(keys, k)
	}
	return keys
}

func (self mapOf[any, T]) Values() []T {
	vals := make([]T, 0, len(self))
	for _, a := range self {
		vals = append(vals, a)
	}
	return vals
}
