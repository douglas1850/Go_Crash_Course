package main

import "fmt"

func main() {
	//pointers used to point to memory address of a value
	a := 5
	b := &a //assign b to a ptr of a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) //%T for type

	//Use * to read val from address
	fmt.Println(*b) //same thing as *&a

	//Change val with ptr
	*b = 10
	fmt.Println(a)
}
