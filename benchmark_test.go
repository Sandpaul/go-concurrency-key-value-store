package main

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

// "time"

func BenchmarkRandInt(b *testing.B) {

	for range b.N {
		rand.Int()
	}
}


func BenchmarkUnbufferedChannel(b *testing.B) {
	Start()

	var wg sync.WaitGroup
	
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			StoreData(strconv.Itoa(i), strconv.Itoa(i))
		}(i)
	}

	wg.Wait()
	
	Stop()
}