package main

import "fmt"

func main() {
	//Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	//FizzBuzz - print out 1 through 100. For every output of 3 output Fizz
	//for every output of 5 output Buzz, and if both output FizzBuzz
	for i := 1; i <= 100; i++ {
		if (i%3) == 0 && (i%5) != 0 {
			fmt.Println("Fizz")
		} else if (i%3) != 0 && (i%5) == 0 {
			fmt.Println("Buzz")
		} else if i%15 == 0 { //if divisible by 15, divisble by 3 and 5
			fmt.Println("Fizz Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
