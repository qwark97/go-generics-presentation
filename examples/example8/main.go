package main

import "fmt"

type XI interface {
	Do() string
	~int | ~string
}

type X string

func (x X) Do() string {
	return string(x)
}

type Y int

func (y Y) Do() string {
	return fmt.Sprint(y)
}

func genericFunction[T XI](arg1, arg2 T) string {
	if arg1 <= arg2 {
		return arg1.Do()
	} else {
		return arg2.Do()
	}
}

func main() {
	x1 := X("asd")
	x2 := X("qwe")
	genericFunction(x1, x2)

	y1 := Y(42)
	y2 := Y(24)
	genericFunction(y1, y2)
}
