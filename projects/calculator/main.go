package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	ret := make(chan int)
	go func(ret chan int) {
		defer close(ret)
		select {
		case f := <-firstChan:
			ret <- f * f
		case s := <-secondChan:
			ret <- s * 3
		case <-stopChan:
			return
		}
	}(ret)
	return ret
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})
	r := calculator(firstChan, secondChan, stopChan)
	firstChan <- 345
	fmt.Println(<-r)
}
