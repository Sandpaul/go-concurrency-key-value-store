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
	go fmt.Printf("[name: %v]\n", FetchData("name"))
	go fmt.Printf("[age: %v]\n", FetchData("age"))
	go fmt.Printf("[city: %v]\n", FetchData("city"))
	go StoreData("email", "paul@example.com")
	go StoreData("phone", "123-456-7890")
	go fmt.Printf("[email: %v]\n", FetchData("email"))
	go fmt.Printf("[phone: %v]\n", FetchData("phone"))
	go StoreData("street", "123 Main St")
	go StoreData("zip", "90001")
	go fmt.Printf("[street: %v]\n", FetchData("street"))
	go fmt.Printf("[zip: %v]\n", FetchData("zip"))
	go StoreData("product_id", "X12345")
	go StoreData("product_name", "Laptop")
	go StoreData("price", "1200.00")
	go fmt.Printf("[product_id: %v]\n", FetchData("product_id"))
	go fmt.Printf("[product_name: %v]\n", FetchData("product_name"))
	go fmt.Printf("[price: %v]\n", FetchData("price"))

}

func main() {

	Start()
	defer Stop()

	simulateConcurrentRequests()
}
