package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var vg = sync.WaitGroup{}
var total = 0

func upNumber() {
	for i := 0; i < 20; i++ {
		x := total
		x++
		time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
		total = x
	}
	vg.Done()
}

func main() {
	vg.Add(2)
	go upNumber()
	go upNumber()
	vg.Wait()
	// expect 40
	fmt.Println(total)
}
