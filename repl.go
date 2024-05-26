package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for{

		fmt.Println("pokedex >")

		scanner.Scan()
		userInput := scanner.Text()
		fmt.Println()
		cleaned := cleanInput(userInput)
		cleanedCommand := cleaned[0]
	// name the map so it can be called
		availableCommands := getCommands(cfg)
// check if command is valid
		commandName,ok := availableCommands[cleanedCommand]

		 if !ok {
			fmt.Println("invalid command")
			continue
		 }
// call anonymous function of the command struct
		 commandName.callback(cfg)
		
		
	
	
	}
}

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name string
	description string
	callback func(cfg *config)
}

// funcition below returns map of command structs with anonymous functions
func getCommands(*config) map[string] cliCommand {
	return map[string] cliCommand { 
		"help":{
			name: "help",
            description: "displays this help menu",
            callback: func(cfg *config) {
                fmt.Println("Welcome to the Pokedex help menu!")
                fmt.Println("Here are the commands you can use:")
                fmt.Println("exit - exits the program")
                fmt.Println("help - displays this help menu")
				fmt.Println("map - displays the next map")
				fmt.Println("mapb - displays the  previous map")
            },
		},
		"exit": {
			name: "exit",
            description: "exits the program",
            callback: func(cfg *config) {
                os.Exit(0)
            },
		
		},
		"map" : {
			name: "map",
            description: "displays the map",
            callback: getPokeLocation,
		},
		"mapb" : {
			name: "mapb",
            description: "displays the map",
            callback: getPrevious,
		},
        
	}
}
