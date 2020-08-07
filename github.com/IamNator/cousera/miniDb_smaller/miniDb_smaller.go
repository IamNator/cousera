package main

import (
	"errors"
	"fmt"
	"strings"
)

//############################
type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (myAnimal Animal) Eat() { fmt.Println(myAnimal.food) }

func (myAnimal Animal) Move() { fmt.Println(myAnimal.locomotion) }

func (myAnimal Animal) Speak() { fmt.Println(myAnimal.sound) }

//################################

func query(userinput []string, animal_list map[string]Animal) error {
	var err error

	animalName := userinput[0]
	animalAttr := userinput[1]

	switch {
	case strings.Contains(animalAttr, "eat"):
		animal_list[animalName].Eat()
	case strings.Contains(animalAttr, "move"):
		animal_list[animalName].Move()
	case strings.Contains(animalAttr, "speak"):
		animal_list[animalName].Speak()
	default:
		err = errors.New("Please enter a valid name (cow|bird|snake) and attribute(eat|speak|move) e.g > cow eat ")
	}
	return err
}

func getUserInput(userinput []string) {
	for i := 0; i < 2; i++ {
		_, err := fmt.Scan(&userinput[i])
		for err != nil {
			fmt.Printf("error :%+v\n", err)
			_, err = fmt.Scan(userinput)
		}
	}
}

func main() {

	AnimalList := make(map[string]Animal)
	AnimalList["cow"] = Animal{food: "grass", locomotion: "walk", sound: "moo"}
	AnimalList["bird"] = Animal{food: "worms", locomotion: "fly", sound: "peep"}
	AnimalList["snake"] = Animal{food: "mice", locomotion: "slither", sound: "hsss"}

	user_input := make([]string, 2)

	fmt.Println("Please enter commands > [cow|bird|snake] [eat|speak|move] :")

	for {
		fmt.Print("> ")

		getUserInput(user_input)

		err := query(user_input, AnimalList)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
