package main

import "fmt"

type Fetch struct {
	key      string
	response chan string
}

func (f *Fetch) Execute(dataStore map[string]string) {
	fmt.Printf("Fetching [%s] from data store...\n", f.key)
	f.response <- dataStore[f.key]
}
