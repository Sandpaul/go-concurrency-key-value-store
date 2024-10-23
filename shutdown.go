package main

import "fmt"

type Shutdown struct{}

func (s *Shutdown) Execute(dataStore map[string]string) {
	fmt.Println("Shutting down")
	close(requests)
}
