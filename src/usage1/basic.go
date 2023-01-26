package generics

func additon[T int | float64](x, y T) T {
	return x + y
}

func use() {
	additon[int](1, 2)
	additon[float64](1.4, 2.6)
}
