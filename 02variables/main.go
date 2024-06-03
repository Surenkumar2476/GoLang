package main

import "fmt"

const LoginToken string = "ajdowjfownfowef" //public - a variable start with capital letter it is a public varibale LoginToken

func main() {

	// string variable
	var username string = "Surenkumar"
	fmt.Println(username)
	fmt.Printf("Variable type is : %T \n", username)

	var islocatefile bool = true
	fmt.Println(islocatefile)
	fmt.Printf("Variable type is : %T \n", islocatefile)

	var smallvalue uint8 = 255
	fmt.Println(smallvalue)
	fmt.Printf("Variable type is : %T \n", smallvalue)

	var smallfloat float32 = 255.34533443434
	fmt.Println(smallfloat)
	fmt.Printf("Variable type is : %T \n", smallfloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is : %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable type is : %T \n", anotherString)

	// implicit types

	var websites = "www.google.com"
	fmt.Println(websites)

	// no var style  // := is the walrus operator    and   walrus is not working in global variable declaration

	numberOfUser := 8978675
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable type is : %T \n", LoginToken)
}
