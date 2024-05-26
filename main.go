package main 

import (
	"fmt"
)

func main() {
	
	command := ""
	for{

		fmt.Println("pokedex >")

		fmt.Scan(&command)
		fmt.Println()

		if command == "help" {
			fmt.Println("Welcome to the Pokedex!\nUsage:\n\nhelp: to display this message\nexit: to exit the application\n")    
        }

		if command == "exit" {
		break	
		}
	}
}