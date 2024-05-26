package main

import (
	
)
type config struct {
	Url string  
	Next *string 
	Previous *string 
}

func main() {
	cfg := config{
		Url: "https://pokeapi.co/api/v2/location-area/",
	}
	startRepl(&cfg)
}