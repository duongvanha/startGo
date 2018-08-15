package main

import (
	"fmt"
	"sync"
	"time"
)

var vg = sync.WaitGroup{}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("foo %d \n", i)
		time.Sleep(time.Millisecond * 500)
	}
	vg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Printf("bar %d \n", i)
		time.Sleep(time.Millisecond * 500)
	}
	vg.Done()
}

func main() {
	vg.Add(2)
	go foo()
	go bar()
	vg.Wait()
}
