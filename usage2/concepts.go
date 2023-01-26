package main

import "fmt"

type magicNumber int

type Number interface {
	~int | int64 | int32 | float32 | float64
}

func GenericFunc[myType rune | Number](value myType) string {
	return fmt.Sprintf("%v", value)
}

func main() {
	var x rune = ' '
	var y float32 = 1.2
	var z magicNumber = 42
	a := GenericFunc[rune](x)
	b := GenericFunc(y)
	c := GenericFunc(z)
	fmt.Printf("%s (%T)\n", a, a)
	fmt.Printf("%s (%T)\n", b, b)
	fmt.Printf("%s (%T)\n", c, c)
}
