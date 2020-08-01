/*
	This program prints on to screen any integer values entered. 
	If a float is entered, it's truncated to an integer

	Author : Nator Verinumbe
	Date   : 8/Aug/2020
	Name   : print number entered as integer
	
*/

/*
	This program prints on to screen any integer values entered. 
	If a float is entered, it's truncated to an integer

	Author : Nator Verinumbe
	Date   : 8/Aug/2020
	Name   : print number entered as integer
	
*/

package main

import "fmt"



func main(){
	fmt.Println(greeting("Nator"))

//	var fruitArr [2] string

	// //Assign Values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	//Declare and assign
	//fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)

}

func greeting(name string) string{
	return "Hello " + name
}