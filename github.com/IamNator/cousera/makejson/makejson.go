
/*
	Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

	Author : Nator Verinumbe
	Date   : 2/Aug/2020
	Name   : makejson.go
	To run : go run makejson.go
	
*/

package main

import (
	"fmt"
	"encoding/json"
)

func main(){


	var name string //for collecting users' name
	var address string //for collecting users' address

	// prompts the user to first enter a name, and then enter an address
	fmt.Print("Please enter your name : ")
	fmt.Scan(&name)
	fmt.Print("Please enter your address : ")
	fmt.Scan(&address)

	var personMap map[string]string //declares a map variable or identifier
	personMap = make(map[string]string) //assigns a map object to the identifier

	personMap["name"] = name
	personMap["address"] = address

	jsonString, _ := json.Marshal(personMap) //converts to json object byte
	

	// fmt.Println("\npersonMap : ")
	// fmt.Println(personMap)           //Prints map object
	fmt.Println("\njsonString : ")
	fmt.Println(string(jsonString))          //Prints json object byte
		
 }