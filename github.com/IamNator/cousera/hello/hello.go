package main

import "fmt"
// type Grades int

// const (
// 	A Grades = iota
// 	B
// 	C
// 	D
// 	F
// )

func main(){
// 		fmt.Printf("%d\n", A)
// 		fmt.Printf("%d\n", B)
// 		fmt.Printf("%d\n", C)
// 		fmt.Printf("%d\n", D)
// 		fmt.Printf("%d\n", F)

// 		say("Nator")

	var appleNum int

	fmt.Printf("Number of apples? \n")

	num, err := fmt.Scan(&appleNum)

	fmt.Printf("appleNum = %d \n", appleNum)
	fmt.Printf("%d\n",num)
	fmt.Printf("%v\n", err)
}


// func say(name string){
// 	switch name {
// 	case "Nator" : 
// 		fmt.Printf("Welcome %s\n", name)
// 	case "Peace" :
// 		fmt.Printf("I love you %s\n", name)
// 	case "Peace Attah":
// 		fmt.Printf("%s can you be my girlfriend\n")
// 	default:
// 		fmt.Printf("I love you\n")
// 	}
// }