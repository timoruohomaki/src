/* INSTRUCTIONS
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the three animals and their associated data which should be hard-coded into your program.

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

// *** FUNCTIONS START HERE ***

func parseInput() (string, string) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	line := scanner.Text()
	line = strings.TrimSpace(line)

	if line == "quit" {
		os.Exit(0)
	}

	lineItems := strings.Split(line, " ")

	return lineItems[0], lineItems[1]
}

func (a Animal) Eat() {
	fmt.Println(a.name, "eats", a.food)
}

func (a Animal) Move() {
	fmt.Println(a.name, "moves by", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.name, "makes noise like", a.noise)
}

// *** MAIN FUNCTION STARTS HERE ***

func main() {

	// setting default values to Animal struct

	var Animals = [...]Animal{
		Animal{"cow", "grass", "walk", "moo"},
		Animal{"bird", "worms", "fly", "peep"},
		Animal{"snake", "mice", "slither", "hsss"},
	}

	fmt.Println("ANIMAL INFORMATION SYSTEM")
	fmt.Println("=========================")
	fmt.Println("Enter search query including animal and function separated by space")
	fmt.Println("Example: cow eat")
	fmt.Println("Exit by entering command: quit")

	// get user input and process request

	for {
		fmt.Print(">")
		animalName, query := parseInput()

		for _, a := range Animals {
			if a.name == animalName {
				switch query {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				}
			}
		}

	}

}
