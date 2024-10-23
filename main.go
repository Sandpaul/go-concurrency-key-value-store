package main

import (
	"fmt"
)

type operation interface {
	execute(dataStore map[string]string)
}

type store struct {
	key      string
	value    string
}

func (s *store) execute(dataStore map[string]string) {
	fmt.Printf("Saving [%s: %s] to data store...\n", s.key, s.value)
	dataStore[s.key] = s.value
	fmt.Println("Data saved successfully.")
}

type fetch struct {
	key 		string
	response 	chan string
}

func (f *fetch) execute(dataStore map[string]string) {
	fmt.Printf("Fetching [%s] from data store...\n", f.key)
	f.response <- dataStore[f.key]
}

type shutdown struct {
	response	chan string
}

func (s *shutdown) execute(dataStore map[string]string) {
	fmt.Println("Shutting down")
	close(requests)
}

var requests chan operation = make(chan operation)
var done chan struct{} = make(chan struct{})

func StoreData(key string, value string) {
	op := &store{key: key, value: value}

	requests <- op
}

func FetchData(key string) string {
	fetchedData := make(chan string)

	op := &fetch{key: key, response: fetchedData}

	requests <- op
	return <-fetchedData
}

func monitorRequests() {
	dataStore := make(map[string]string)

	for op := range requests {
		op.execute(dataStore)
	}

	fmt.Println("All requests processed")
	close(done)
}

func Start() {
	go monitorRequests()
}

func Stop() {
	shutdown := &shutdown{response: nil}
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
