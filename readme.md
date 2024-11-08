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
main.go
```go
package main

import (
	"errors"
	"log"
)

func main() {
	result, err := devide(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result of division is", result)
}

func devide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot devide by 0")
	}

	result = x / y
	return result, nil
}
```

main_test.go
```go
package main

import "testing"

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-date", 100.0, 10.0, 10.0, false},
	{"invalid-date", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := devide(tt.divident, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not got one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but gone one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f bit got %f", tt.expected, got)
		}
	}
}

func TestDevide(t *testing.T) {
	_, err := devide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDevide(t *testing.T) {
	_, err := devide(10.0, 0)
	if err == nil {
		t.Error("Dig not get an error when we should not have")
	}
}

```

Command for running test:
- go test
- go test -v
- go test -cover
- go test -coverprofile=coverage.out && go tool cover -html=coverage.out


## Section 3: Building a Basic Web Application
### Making a "Hello, World" web application
main.go
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
```

### Making our application module-ready
go mod init _MODULE_NAME_