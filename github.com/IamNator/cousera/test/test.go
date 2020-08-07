package main

import (
	"fmt"
)

var animals = map[string]Animal{
	"cow": {
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	},
	"bird": {
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	},
	"snake": {
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	},
}

const unknownAnimalMessage = "unknown animal, it should be either “cow”, “bird”, or “snake”"
const unknownActionMessage = "unknown action, it should be either “eat”, “move”, or “speak”"

var UnknownAnimal = Animal{
	food:       unknownAnimalMessage,
	locomotion: unknownAnimalMessage,
	noise:      unknownAnimalMessage,
}

func main() {
	for {
		fmt.Print("> ")

		var animalType, action string
		_, err := fmt.Scanf("%s %s", &animalType, &action)
		if err != nil {
			fmt.Println("Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.\n" + err.Error())
			continue
		}

		animal := animalFor(animalType)

		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println(unknownActionMessage)
		}

	}
}

func animalFor(animalType string) Animal {
	if foundAnimal, ok := animals[animalType]; ok {
		return foundAnimal
	}
	return UnknownAnimal
}

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}
