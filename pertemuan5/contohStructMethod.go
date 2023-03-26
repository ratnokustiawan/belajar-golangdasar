package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) sayHello() {
	fmt.Printf("Hello, my name is %s\\n", p.name)
}

func main() {
	p := person{name: "John", age: 30}
	p.sayHello()
}
