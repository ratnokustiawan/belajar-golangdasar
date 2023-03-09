package main

import (
	"fmt"
	"math"
)

func main() {
	a := 16.0

	squareRoot := math.Sqrt(a)
	fmt.Println("squareRoot= ", squareRoot)

	sin := math.Sin(a)
	fmt.Println("sin= ", sin)

	cos := math.Cos(a)
	fmt.Println("cos= ", cos)

	tan := math.Tan(a)
	fmt.Println("tan= ", tan)

	log := math.Log(a)
	fmt.Println("log= ", log)

	exp := math.Exp(a)
	fmt.Println("exp= ", exp)

	round := math.Round(a)
	fmt.Println("round= ", round)

	abs := math.Abs(a)
	fmt.Println("abs= ", abs)

	pow := math.Pow(a, 2)
	fmt.Println("pow= ", pow)
}
