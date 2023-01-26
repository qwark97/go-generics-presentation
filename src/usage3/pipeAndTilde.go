package main

import "fmt"

type CoolNumber int
type CoolerNumber CoolNumber

type Number interface {
	~int | int64 | int32 | float32 | float64
}

func GenericFunc[myType rune | Number](value myType) string {
	return fmt.Sprintf("%v", value)
}
