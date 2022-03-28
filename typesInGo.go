package main

import "fmt"

//Insert variables declarations here based on activity

func tellMeTypes() {
	//insert code here to print out types of variables
	text := "The following is the account information."
	firstName := "Luke"

	fmt.Printf("%T %T", text, firstName)
}

func main() {
	tellMeTypes()
}
