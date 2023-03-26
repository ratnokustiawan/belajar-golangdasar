package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) ChangeName(name string) {
	p.Name = name
}

func main() {
	person := Person{Name: "John", Age: 30}
	fmt.Println(person.Name) // Output: John
	person.ChangeName("Doe")
	fmt.Println(person.Name) // Output: Doe
}
