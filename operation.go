package main

type Operation interface {
	Execute(dataStore map[string]string)
}
