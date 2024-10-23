package main

import (
	"testing"
)

func TestStoreExecute(t *testing.T) {

	dataStore := make(map[string]string)

	store := Store{key: "name", value: "paul"}
	store.Execute(dataStore)

	expected := "paul"
	actual := dataStore["name"]

	if actual != expected {
		t.Errorf("Expected: %s, actual: %s", expected, actual)
	}
}

func TestFetchExecute(t *testing.T) {

	dataStore := make(map[string]string)
	dataStore["age"] = "34"

	responseChan := make(chan string)
	fetch := Fetch{key: "age", response: responseChan}

	go fetch.Execute(dataStore)

	expected := "34"
	actual := <- responseChan

	if actual != expected {
		t.Errorf("Expected: %s, actual: %s", expected, actual)
	}
}

func TestIntegrationTest(t *testing.T) {

	Start()
	defer Stop()

	StoreData("name", "ralph")
	
	expected := "ralph"
	actual := FetchData("name")

	if actual != expected {
		t.Errorf("Expected: %s, actual: %s", expected, actual)
	}

	if result := FetchData("non existent key"); result != "" {
		t.Errorf("Expected '', got %s", result)
	}
}
