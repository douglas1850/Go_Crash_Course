package main

import "fmt" //formater. Always use double quotes

func main() {
	//Define map
	//name for keys, emails for values
	emails := make(map[string]string)
	//first string is for key, second is for value

	//Assign key values
	emails["Bob Builder"] = "bob@gmail.com"
	emails["Mike Situation"] = "theSitch@gmail.com"
	emails["PaulyD"] = "paulyd@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob Builder"]) //will output the value for the key
	fmt.Println(len(emails))

	//Delete from map
	delete(emails, "PaulyD") //delete entry for PD
	fmt.Println(emails)
	fmt.Println(len(emails))

	//Declare map and add key values
	address := map[string]int{"Bob": 1231, "Mike": 5654}
	fmt.Println(address)
}
