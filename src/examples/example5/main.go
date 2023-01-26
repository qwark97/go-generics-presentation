package main

import "errors"

type typeSet interface {
	error | any
}

func genericFunction[T typeSet](arg T) T {
	return arg
}

func main() {
	err := errors.New("")
	genericFunction(err)
}
