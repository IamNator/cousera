package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Printf("Enter string\n")

	var userInput string = ""
	/*num,err := */fmt.Scan(&userInput)
	
	if strings.ContainsAny(userInput, "ian" ) {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}