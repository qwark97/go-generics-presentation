package main

import "fmt"

type someFunction func()

type someOtherFunction func(arg string)

func genericFunction[T someFunction | someOtherFunction](passedArgument T) error {
	passedArgument()
	return nil
}

func main() {
	f1 := func() {
		fmt.Println("f1 magic")
	}
	f2 := func(arg string) {
		fmt.Println("f2 magic:", arg)
	}

	genericFunction[someFunction](f1)
	genericFunction[someOtherFunction](f2)
}
