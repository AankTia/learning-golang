# Building Modern Web Applications with Go (Golang)

## Section 1: Introduction
### Why Go?

## Section 2: Overview of the Go Language
### Variable & Function
### Pointers
### Types and Struct
### Receivers:Struct with functions
### Other Data Structures: Maps and Slice
### Decision Structures
### Loops and rangin over data
### Interfaces
### Packages
### Channels
### Reading and Writting JSON
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasCat bool `json:"has_cat"`
}

func main() {
	myJson := `
[
	{ 
		"first_name": "Tia",
		"last_name": "Widi",
		"hair_collor": "black",
		"has_cat": false
	},
	{
		"first_name": "Aradea",
		"last_name": "Widi",
		"hair_collor": "black",
		"has_cat": false
	}
]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Println("unmarshalled: %w", unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Ardhana"
	m1.LastName = "Widi"
	m1.HairColor = "black"
	m1.HasCat = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m1.FirstName = "Fanny"
	m1.LastName = "Widi"
	m1.HairColor = "black"
	m1.HasCat = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}
```
### Writing Tests in Go