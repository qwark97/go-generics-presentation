package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func getSmallest[T constraints.Ordered](values ...T) (T, error) {
	var smallest T
	if len(values) == 0 {
		return smallest, fmt.Errorf("not enough values")
	}

	smallest = values[0]
	for _, x := range values {
		if x < smallest {
			smallest = x
		}
	}
	return smallest, nil
}

func main() {
	strings := []string{"q", "b", "qwe", "asd"}
	ints := []int{8, 42, 3, 10}

	fmt.Println(getSmallest(strings...))
	fmt.Println(getSmallest(ints...))
}
