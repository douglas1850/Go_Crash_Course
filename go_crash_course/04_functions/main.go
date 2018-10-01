package main

import "fmt"

func greeting(name string) string {
	//string return type
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Douglas"), getSum(1234, 616))
}
