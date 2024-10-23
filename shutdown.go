package main

import "fmt"

type Shutdown struct {
	response chan string
}

func (s *Shutdown) Execute(dataStore map[string]string) {
	fmt.Println("Shutting down")
	close(requests)
}
