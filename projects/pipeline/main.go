package main

import "fmt"

func removeDuplicates(in, out chan string) {
	go func(in, out chan string) {
		temp := ""
		for {
			select {
			case val, isOpen := <-in:
				if isOpen {
					if val != temp {
						temp = val
						out <- val
					}
				} else {
					close(out)
					return
				}
			}
		}
	}(in, out)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)

		for _, r := range "111112334456" {
			inputStream <- string(r)
		}
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
	fmt.Println()
}
