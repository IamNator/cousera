/*
	Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.



	Author    : Nator Verinumbe
	Date   	  : 3/Aug/2020
	Name   	  : bubblesort.go
	How to run: go run bubblesort.go
*/

package main

import (
	"fmt"
)

func main() {
	var userInput [10]int //for storing user inputs (array)

	fmt.Println("Please enter 10 intergers pressing enter between each entry") //Prompts user for input
	for i, _ := range userInput {                                              //gets sequence of 10 integers from user
		_, err := fmt.Scan(&userInput[i])
		if err != nil {
			fmt.Println("Please enter an integer ")
		}
	}

	userInputSlice := userInput[0:10] //creates a slice from array userInput
	BubbleSort(userInputSlice)        //sorts slice using bubble sort algorithm

	fmt.Println("Sorted slice : ")
	fmt.Println(userInputSlice) //Prints sorted slice
}

func BubbleSort(data []int) { //performs a bubble sort
	for j := len(data) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if data[i] > data[i+1] {
				Swap(data, i)
			}
		}
	}
}

func Swap(intSlice []int, index int) { //swaps two values
	intSlice[index], intSlice[index+1] = intSlice[index+1], intSlice[index]
}
