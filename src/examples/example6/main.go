package main

type X struct{}

func (x X) Do()   {}
func (x X) Did()  {}
func (x X) Done() {}

type Y struct{}

func (y Y) Do()  {}
func (y Y) Did() {}

type Z struct{}

func (z Z) Do()   {}
func (z Z) Done() {}

type typeSet interface {
	X | Y | Z
}

func genericFunction[T typeSet](arg T) T {
	return arg
}

func main() {
	genericFunction(X{})
	genericFunction(Y{})
	genericFunction(Z{})
}
