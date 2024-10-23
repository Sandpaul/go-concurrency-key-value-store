package main

import (
	"fmt"
	"time"
)

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

func StoreData(key string, value string) {
	op := operation{action: "store", key: key, value: value}

	requests <- op
}

func FetchData(key string) string {
	fetchedData := make(chan string)

	op := operation{action: "fetch", key: key, value: "", response: fetchedData}

	requests <- op
	return <-fetchedData
}

func monitorRequests() {
	dataStore := make(map[string]string)
	
	for op := range requests {
		switch op.action {
		case "store":
			fmt.Printf("Saving %s: %s to data store...\n", op.key, op.value)
			dataStore[op.key] = op.value
			fmt.Println("Data saved successfully.")
		case "fetch":
			fmt.Printf("Fetching %s from data store...\n", op.key)
			op.response <- dataStore[op.key]
		}
	}
}

func Start() {
	go monitorRequests()
}

func main() {
	Start()

	go StoreData("name", "paul")
	go StoreData("age", "34")
	go fmt.Println("name:" + FetchData("name"))
	go fmt.Println("age:" + FetchData("age"))

	time.Sleep(15 * time.Second)
}