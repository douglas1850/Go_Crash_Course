package main

import "fmt"

//anonymous functions are closures. Define function inline without having to name it

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i)) //adds next number in the loop
	}
}
