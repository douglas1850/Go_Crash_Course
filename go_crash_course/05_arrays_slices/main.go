package main

import "fmt"

func main() {
	//Arrays must be a fixed length and name the types
	//Slices are arrays without a fixed type

	var fruitArr [2]string

	//assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])

	//declare and assign
	fruitArr2 := [2]string{"Bannana", "Pinneapple"}
	fmt.Println(fruitArr2)

	//Slices
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3]) //starts at 1 and stops before 3

}
