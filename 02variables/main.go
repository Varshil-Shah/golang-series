package main

import "fmt"

// GoLang syntax is not allowed globally
// JwtToken := "ndsjdvnkjvbeujbverjsbv"
var JwtToken = "djvfjvjbjdbc jd " // Starting letter is Capital then Public else Private

func main() {
	/*
		DEFAULT VALUES -
		int (0)
		string ('')
		bool (false)
	*/

	var username string = "varshil"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallNumber byte = 255
	fmt.Println(smallNumber)
	fmt.Printf("Variable is of type: %T \n", smallNumber)

	var smallFloat float32 = 5152.562348762234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 5152.562348762234
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	// Default values of an variables
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Default value of type: %T \n", anotherVariable)

	var emptyString string
	fmt.Println(emptyString)
	fmt.Printf("Default value of type: %T \n", emptyString)

	// Implicit type
	var numberOfUsers = 500
	fmt.Println(numberOfUsers)

	// No var style
	emailAddress := "varshil@gmail.com"
	fmt.Println(emailAddress)

	reviews := 4.8
	fmt.Println(reviews)
}
