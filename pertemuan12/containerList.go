package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()

	myList.PushBack("Golang")
	myList.PushBack("Python")
	myList.PushBack("Java")
	myList.PushBack("PHP")

	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
