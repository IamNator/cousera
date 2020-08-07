package main

import (
	"fmt"
	"math"
)

const (
	acceleration        = iota //0
	initialVelocity            //1
	initialDisplacement        //2
)

func main() {

	//prompts the user to enter values for acceleration, initial velocity, and initial displacement
	fmt.Println("Please enter values for acceleration, initial velocity, and initial displacement")
	car := make([]float64, 3) //for user input
	getUserInput(car)

	findDisplacement := GenDisplaceFn(car)

	var time_byUser float64
	fmt.Print("Please enter time : ")
	fmt.Scan(&time_byUser)

	fmt.Print("Hence displacement : ")
	fmt.Printf("%.3f  meters ", findDisplacement(time_byUser))
}

func getUserInput(user_input []float64) {

	for i, _ := range user_input {

		_, err := fmt.Scan(&user_input[i])

		for err != nil { //Loops back when input is not a number till a number is entered
			fmt.Println("Please enter a number ")
			_, err = fmt.Scan(&user_input[i])
		}

	}

}

func GenDisplaceFn(sensor []float64) func(float64) float64 {

	return func(time_here float64) float64 {

		displacement := 0.5*sensor[acceleration]*math.Pow(time_here, 2) + sensor[initialVelocity]*time_here + sensor[initialDisplacement]

		return displacement
	}
}
