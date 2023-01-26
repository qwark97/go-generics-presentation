package main

import "fmt"

type ExpectedReturnTypes interface {
	uint | bool
}

type ExpectedArgsTypes interface {
	int | string
}

func decorator[expectedFunction func(x expectedArg) expectedReturn, expectedReturn ExpectedReturnTypes, expectedArg ExpectedArgsTypes](passedFunction expectedFunction, passedArgument expectedArg) expectedReturn {
	magicLogicDecoratorDoes()
	return passedFunction(passedArgument)
}

func main() {
	f1 := func(x string) uint {
		fmt.Printf("¯\\_(%s%s\n", "ツ", x)
		return 0
	}
	f2 := func(x int) bool {
		fmt.Println("answer:", x)
		return true
	}

	decorator(f2, 42)
	decorator(f1, ")_/¯")
}

func magicLogicDecoratorDoes() {
	/*
		...
	*/
}
