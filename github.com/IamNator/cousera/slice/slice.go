
/*
	Module 3 Activity: slice.go
	This program sorts and prints out integer inputs

	Author : Nator Verinumbe
	Date   : 2/Aug/2020
	Name   : sort integer values
	
*/

package main

import (
	"fmt"  //for interacting with the console 
	"sort" //for sorting
	"strconv"  //for string type conversion
)

func main(){

	numSlice := make([]int , 3)  //creates and initialize to zero a slice of length 3

	var userInput string  //takes user inputs
	var counts int = 0   //to keep track of the number of interation of main loop

	for userInput != "X" {
		fmt.Println("Please enter an integer or enter X to close the program :")
		fmt.Scan(&userInput)

		if userInput != "X" {
			userInputnum, err := strconv.Atoi(userInput) //converts input to integers
			

			if err == nil {     //check if it is an integer value entered
			
				if counts <= 2 {     //stores first 3 inputs
					if numSlice[0] == userInputnum {    //for negative number inputs
						numSlice[counts] = userInputnum
					} else {
						numSlice[0] = userInputnum
					}					
				} else {			//appends if the length of numSlice is exceeded
					numSlice = append(numSlice, userInputnum)  
				}
				
				sort.Ints(numSlice) //Sorts the array

				fmt.Println(numSlice) 
				fmt.Println(err) //Prints error message

			counts++  //increments counts  --- warning this can cause memory leaks if the program continually runs
			} else {
				fmt.Println("Please enter an integer") 
			}
		}
		
	}

}