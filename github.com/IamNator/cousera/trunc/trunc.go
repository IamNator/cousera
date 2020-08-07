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

	var s []int 
	PrintSlice(s)

	s = append(s, 0)
	PrintSlice(s)

	s = append(s, 1, 2)
	PrintSlice(s)

}

func PrintSlice(s []int){
	fmt.Printf("%d  %d  %v\n", len(s), cap(s), s)
}