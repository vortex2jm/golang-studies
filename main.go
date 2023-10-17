package main

import (
	"fmt"
	math "go-studies/my-math" // Assigning another name to the package
)

func main() {
	print("Hello world\n")
	fmt.Println("Hello world again")

	// Declaring and assigning at the same time
	a := "Joao"
	b := 10
	c := 20.5
	d := true

	println()
	fmt.Printf("%T - Variable type 'a'\n", a)
	fmt.Printf("%v - Variable value 'a'\n", a)
	fmt.Printf("%T - Variable type 'b'\n", b)
	fmt.Printf("%v - Variable value 'b'\n", a)
	fmt.Printf("%T - Variable type 'c'\n", c)
	fmt.Printf("%v - Variable value 'c'\n", a)
	fmt.Printf("%T - Variable type 'd'\n", d)
	fmt.Printf("%v - Variable value 'd'\n", a)

	//Declaring and assigning separately
	var e int
	e = 13

	//Accessing function from another path
	f := math.Sum(b, e)

	fmt.Println("Valor de f =", f)
}
