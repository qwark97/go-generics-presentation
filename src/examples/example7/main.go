package main

type XI interface {
	Did()
}
type YI interface {
	Did()
	Done()
}
type ZI interface {
	Do()
	Did()
}

type typeSet interface {
	XI | YI | ZI
}

type X struct {
}

func (x X) Did() {}

func genericFunction[T typeSet](arg T) T {
	return arg
}

func main() {
	genericFunction(X{})
}
