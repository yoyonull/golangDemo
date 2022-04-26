package main

import (
	"fmt"
	"sync"
)

func main() {

	str, num := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-num:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				str <- true
			}

		}
	}()

	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-str:
				if i >= 'Z' {
					wait.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				num <- true

			}
		}
	}(&wait)

	num <- true
	wait.Wait()
	fmt.Println()
}
