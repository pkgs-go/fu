package fu

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a ...T) T {
	_, v := MinIndex(a...)
	return v
}

func MinOr[T constraints.Ordered](dflt T, a ...T) T {
	if len(a) > 0 {
		return Min(a...)
	}
	return dflt
}

func MinOf[T constraints.Ordered](v T, a ...T) T {
	if len(a) > 0 {
		if x := Min(a...); x < v {
			return x
		}
	}
	return v
}

func MinIndex[T constraints.Ordered](a ...T) (int, T) {
	j, y := 0, a[0]
	for i, x := range a[1:] {
		if x < y {
			j, y = i, x
		}
	}
	return j, y
}

func Max[T constraints.Ordered](a ...T) T {
	_, v := MaxIndex(a...)
	return v
}

func MaxOr[T constraints.Ordered](dflt T, a ...T) T {
	if len(a) > 0 {
		return Max(a...)
	}
	return dflt
}

func MaxOf[T constraints.Ordered](v T, a ...T) T {
	if len(a) > 0 {
		if x := Max(a...); x > v {
			return x
		}
	}
	return v
}

func MaxIndex[T constraints.Ordered](a ...T) (int, T) {
	j, y := 0, a[0]
	for i, x := range a[1:] {
		if x > y {
			j, y = i, x
		}
	}
	return j, y
}
