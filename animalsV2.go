package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Name() string
	Eat()
	Move()
	Speak()
}

type cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

type bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

type snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

// *** FUNCTIONS START HERE ***

func parseInput() (string, string, string) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	line := scanner.Text()
	line = strings.TrimSpace(line)

	// quit handled here because it is not a command sequence

	if line == "quit" {
		os.Exit(0)
	}

	lineItems := strings.Split(line, " ")

	return lineItems[0], lineItems[1], lineItems[2]
}

func (a cow) Name() string {
	return a.name
}

func (a bird) Name() string {
	return a.name
}

func (a snake) Name() string {
	return a.name
}

func (a cow) Eat() {
	fmt.Println(a.name, "eats", a.food)
}

func (a bird) Eat() {
	fmt.Println(a.name, "eats", a.food)
}

func (a snake) Eat() {
	fmt.Println(a.name, "eats", a.food)
}

func (a cow) Move() {
	fmt.Println(a.name, "moves by", a.locomotion)
}

func (a bird) Move() {
	fmt.Println(a.name, "moves by", a.locomotion)
}

func (a snake) Move() {
	fmt.Println(a.name, "moves by", a.locomotion)
}

func (a cow) Speak() {
	fmt.Println(a.name, "makes noise like", a.noise)
}

func (a bird) Speak() {
	fmt.Println(a.name, "makes noise like", a.noise)
}

func (a snake) Speak() {
	fmt.Println(a.name, "makes noise like", a.noise)
}

// *** MAIN FUNCTION STARTS HERE ***

func main() {

	animals := make(map[string]Animal, 0)

	fmt.Println("THE ANIMAL INFORMATION SYSTEM")
	fmt.Println("=============================")
	fmt.Println("Available commands:")
	fmt.Println("newanimal query quit")
	fmt.Println("parameters for newanimal:")
	fmt.Println("name type(cow|bird|snake)")
	fmt.Println("parameters for query:")
	fmt.Println("name function(eat|move|speak)")

	// get user input and process request

	for {
		fmt.Print(">")
		cmd, name, query := parseInput()

		switch cmd {
		case "newanimal":
			switch query {
			case "cow":
				animals[name] = cow{name, "grass", "walk", "moo"}
			case "bird":
				animals[name] = bird{name, "worms", "fly", "peep"}
			case "snake":
				animals[name] = snake{name, "mice", "slither", "hsss"}
			}

			fmt.Println("Created it!")

		case "query":
			a, err := animals[name]
			if a == nil {
				fmt.Println("ERROR: Not found, message:", err)
				continue
			}
			switch query {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			}

		case "quit":
			os.Exit(0)
		}

	}

}
