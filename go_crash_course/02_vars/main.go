package main

import (
	"fmt"
)

//must always use vars. Unused vars give error
//can declare variables up here as globals

func main() {
	//string
	//bool
	//int int8 int16 int32 int64
	//uint unsigned ints, no negative nums also goes to 64
	//byte alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128 - really large numbers

	//using var
	var name string = "Douglas O'Meara" //doesn't actually need 'string', it's inferred
	var age = 25
	const isSomething = true //const cannot be redefined bcuz it's a Constant

	//shorthand
	nickName := "Doug"
	size := 1.3

	example1, example2 := "some text", "some more text"

	fmt.Println(name, age, nickName)
	fmt.Printf("%T %T\n", isSomething, size) //gives type of variable
	fmt.Println(example1, example2)

}
