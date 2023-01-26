package main

import "fmt"

func genericFunction[T any](passedArgument []T) {
	fmt.Printf("%v, %T\n", passedArgument, passedArgument)

	for _, x := range passedArgument {
		fmt.Printf("%v, %T\n", x, x)
	}
}

func main() {
	x := []int{1, 2, 3}
	y := []string{"a", "b", "c"}

	genericFunction(x)
	genericFunction(y)
}
