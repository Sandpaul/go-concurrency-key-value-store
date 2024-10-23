package main

import "fmt"

// Create a key-value store:
//// Store(string key, string value)
//// Fetch(string key) -> returns the value
//// Start()
//// Stop()

// Apply the Actor model:
//// As per the slides
//// Protected resource: a map holding the data

// Constraints:
//// Cannot use sync.Map
//// Cannot use sync.Mutex

type operation struct {
	action string
	key string
	value string
	response chan string
}

var requests chan operation = make(chan operation)

func Store(key string, value string) {
	storeRequest := operation{action: "store", key: key, value: value}

	requests <- storeRequest
}


func main() {
	fmt.Println("Hello, World!")
}