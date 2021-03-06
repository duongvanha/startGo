package main

import "fmt"

func main() {

	n := 10
	c := make(chan int)

	done := make(chan bool, n)

	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for u := range c {
		fmt.Println(u)
	}
}
