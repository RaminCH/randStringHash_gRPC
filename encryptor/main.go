package main

import (
	"crypto/sha256"
	"fmt"
	"sync"
	"time"

	a "github.com/RaminCH_self/Gransoft/v2/randomizer"
)

// NewSHA256 ...
func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

//Combine func ...
func Combine(n []string) []byte {
	var i int = 4
	str := a.ListOfStrings(a.RandomString(i), i)
	var x = []byte{}
	for i := 0; i < len(str); i++ {
		b := []byte(str[i])
		for j := 0; j < len(b); j++ {
			x = append(x, b[j])
		}
	}
	return NewSHA256(x)
}

//Process function
func Process(i int, b []byte, wg *sync.WaitGroup) {
	go Combine(a.ListOfStrings(a.RandomString(i), i))
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	//pool
	var wg sync.WaitGroup
	num := 4
	var b = Combine(a.ListOfStrings(a.RandomString(4), 4))

	for i := 0; i < num; i++ {
		wg.Add(1)
		go Process(i, b, &wg)
	}
	wg.Wait()
	fmt.Println(string(b))
	fmt.Println("All go routines finished executing")
}
