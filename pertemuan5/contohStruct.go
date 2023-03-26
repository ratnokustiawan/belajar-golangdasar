package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var person1 Person
	person1.Name = "John Doe"
	person1.Age = 30

	fmt.Println(person1.Name)
	fmt.Println(person1.Age)
}
