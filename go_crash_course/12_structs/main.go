package main

import (
	"fmt"
	"strconv"
)

//strconv is string converter

//structs: like classes. No classes in Go.
//Create struct for a person and give them values like age, email
//types of structs: pointer recievers and value recievers

//Define person struct
type Person struct {
	firstName string
	lastName  string
	age       int
	gender    string

	//can put same datatypes on same line
	//ex: firstName, lastName string
}

//Greeting method (value reciever)
func (p Person) greet() string { //p is identifier - can be anything. greet() is name. string return value
	return "Hello, my name is " + p.firstName + " and I'm " + strconv.Itoa(p.age) + " years old\n"
} //doesn't change any values, just uses existing ones to create a greeting

//hasBirthday method (pointer reciever) - we're going to change something
func (p *Person) hasBirthday() { //not returning anything, so no return value
	p.age++
}

//getsMarried method (pointer reciever) - change last name dependent on gender
func (p *Person) getsMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//Init person using Struct
	person1 := Person{firstName: "Bob", lastName: "the Builder", age: 32, gender: "Male"}
	fmt.Println(person1)

	//Alternative
	person2 := Person{"Sharon", "FeminineLastName", 32, "Female"}
	fmt.Println(person2)

	//Get single field
	fmt.Println(person1.firstName)

	//Alter field
	person1.age++
	fmt.Println(person1.age)

	//Call the greet method
	fmt.Println(person1.greet())

	//Call hasBirthday method
	person1.hasBirthday()
	fmt.Println(person1.age)

	fmt.Println(person2.lastName)
	person2.getsMarried("the Builder")
	fmt.Println(person2.lastName)
}
