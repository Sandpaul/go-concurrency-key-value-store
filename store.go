package main

import "fmt"

type Store struct {
	key   string
	value string
}

func (s *Store) Execute(dataStore map[string]string) {
	fmt.Printf("Saving [%s: %s] to data store...\n", s.key, s.value)
	dataStore[s.key] = s.value
	fmt.Println("Data saved successfully.")
}
