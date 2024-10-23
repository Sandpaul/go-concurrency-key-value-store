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

func CreateStoreRequest(key string, value string) {
	storeRequest := operation{action: "store", key: key, value: value}

	requests <- storeRequest
}

func CreateFetchRequest(key string) string {
	fetchedData := make(chan string)

	op := operation{action: "fetch", key: key, value: ""}

	requests <- op
	return <- fetchedData
}

func Start() {
	go monitorRequests()
}

func monitorRequests() {
	dataStore := make(map[string]string)

	for op := range requests {
		switch op.action {
		case "store":
			fmt.Printf("Saving %s: %s to data store...", op.key, op.value)
			dataStore[op.key] = op.value
			fmt.Println("Data saved successfully.")
		case "fetch":
			fmt.Printf("Fetching %s from data store...", op.key)
			op.response <- dataStore[op.key]
		}
	}
}



func main() {
	fmt.Println("Hello, World!")
}