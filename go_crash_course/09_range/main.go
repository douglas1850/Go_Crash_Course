package main

import "fmt" //formater. Always use double quotes

func main() {
	ids := []int{33, 34, 54, 56, 98, 1} //slice of ids

	//Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum = ", sum)

	//Range with map
	address := map[string]int{"Bob": 1231, "Mike": 5654}

	for i, j := range address {
		fmt.Printf("%s: %d\n", i, j) //i has name, j has address
	}

	for k := range address {
		fmt.Println("Name: " + k)
	}
}
