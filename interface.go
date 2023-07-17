package main

import "fmt"

type phone interface {
	call()
}
type NokiaPhone struct {
}

func (nokiaphone NokiaPhone) call() {
	fmt.Println("i am nokia,i can call you")
}

func main() {
	var phone phone
	phone = new(NokiaPhone)
	phone.call()
}
