package main

import "fmt"

// first letter capital maane public
const LoginToken string = "hdbfe"

func main() {

	var username string = "Yash"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isBool bool = true
	fmt.Println(isBool)
	fmt.Printf("Variable is of type: %T \n", isBool)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// 5 decimals
	var smallFloat float32 = 255.767862
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// more control 13 decmals
	var bigFloat float32 = 255.76786273981274379469
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type

	var website = "yashdhasmana.hashnode.dev"
	fmt.Println(website)

	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
