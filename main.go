package main

import (
	"fmt"
)

var requests chan Operation = make(chan Operation)
var done chan struct{} = make(chan struct{})

func StoreData(key string, value string) {
	op := &Store{key: key, value: value}

	requests <- op
}

func FetchData(key string) string {
	fetchedData := make(chan string)

	op := &Fetch{key: key, response: fetchedData}

	requests <- op
	return <-fetchedData
}

func monitorRequests() {
	dataStore := make(map[string]string)

	for op := range requests {
		op.Execute(dataStore)
	}

	fmt.Println("All requests processed")
	close(done)
}

func Start() {
	go monitorRequests()
}

func Stop() {
	shutdown := &Shutdown{}
	requests <- shutdown
	<-done
}

func simulateConcurrentRequests() {
	go StoreData("name", "paul")
	go StoreData("age", "34")
	go StoreData("city", "New York")
	go fmt.Println("name: " + FetchData("name"))
	go fmt.Println("age: " + FetchData("age"))
	go fmt.Println("city: " + FetchData("city"))
	go StoreData("email", "paul@example.com")
	go StoreData("phone", "123-456-7890")
	go fmt.Println("email: " + FetchData("email"))
	go fmt.Println("phone: " + FetchData("phone"))
	go StoreData("street", "123 Main St")
	go StoreData("zip", "90001")
	go fmt.Println("street: " + FetchData("street"))
	go fmt.Println("city: " + FetchData("city"))
	go fmt.Println("zip: " + FetchData("zip"))
	go StoreData("product_id", "X12345")
	go StoreData("product_name", "Laptop")
	go StoreData("price", "1200.00")
	go fmt.Println("product_id: " + FetchData("product_id"))
	go fmt.Println("product_name: " + FetchData("product_name"))
	go fmt.Println("price: " + FetchData("price"))
}

func main() {

	Start()
	defer Stop()

	simulateConcurrentRequests()
}
