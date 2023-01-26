package main

type X struct {
	name string
}

type Y struct {
	name string
}

type typeSet interface {
	X | Y | int
}

func max[T typeSet](a, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	max(1, 2)
}
