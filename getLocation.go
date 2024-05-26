package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "time"
)
var baseUrl = "https://pokeapi.co/api/v2/location-area/"


func getPokeLocation(cfg *config) {
	if cfg.Next == nil {
	res, err := http.Get(baseUrl)
    if err!= nil {
        panic(err)
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err!= nil {
        panic(err)
    }
	locData := LocationArea{}
	json.Unmarshal(body, &locData)
	
    for location := range locData.Results {
		fmt.Println(locData.Results[location].Name)
	}
	cfg.Next = locData.Next
	cfg.Previous = locData.Previous
}else{
	res, err := http.Get(*cfg.Next)
    if err!= nil {
        panic(err)
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err!= nil {
        panic(err)
    }
    locData := LocationArea{}
    json.Unmarshal(body, &locData)
    
    for location := range locData.Results {
        fmt.Println(locData.Results[location].Name)
    }
    cfg.Next = locData.Next
    cfg.Previous = locData.Previous
}
}

func getPrevious(cfg *config) {
	res, err := http.Get(*cfg.Previous)
    if err!= nil {
        panic(err)
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err!= nil {
        panic(err)
    }
    locData := LocationArea{}
    json.Unmarshal(body, &locData)
    
    for location := range locData.Results {
        fmt.Println(locData.Results[location].Name)
    }
    cfg.Next = locData.Next
    cfg.Previous = locData.Previous
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string   `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

