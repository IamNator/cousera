/*
	Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

	Author : Nator Verinumbe
	Date   : 2/Aug/2020
	Name   : read.go
	To run : go run read.go
	
*/


package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){

	type name struct{
		fname string
		lname string 
	}
	
	
	var filename string
	
	fmt.Print("Please the file name plus extension e.g. test.txt : ") //Prompts user for file name input
	fmt.Scan(&filename)                                               //gets file name from user
	readFile, err := os.Open(filename) 								  //returns a pointer to text file
	if err != nil {                                                   // Prints error if any
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)   //creates a scanner for the text file
	fileScanner.Split(bufio.ScanLines)          //to scan for lines in text file


	var nameSlice []name     //Slice of name structs


	for fileScanner.Scan() {                                                   //goes through the text file line by line
		wordScanner := bufio.NewScanner(strings.NewReader(fileScanner.Text())) //creates a scanner for each line of text from text file
		wordScanner.Split(bufio.ScanWords)                                     //to scan for words
		var p1 name
		i := 0
		for wordScanner.Scan() {                                                                  //goes through the line word by word
			if i==0 {p1.fname = wordScanner.Text()} else if i==1 {p1.lname = wordScanner.Text()}  //write fname and lname
			i++
		}
		nameSlice = append(nameSlice, p1)														 //appends struct to slice
	}

	readFile.Close()

	for _, names := range nameSlice {              //iterates through nameSlice and prints the first and last names found in each struct
		fmt.Println(names.fname , names.lname)     //prints first and last name in each struct
	}

	

}