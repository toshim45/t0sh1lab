package main

import (
	"log"
	"sync"
)

func main() {
	log.Println("[main] hi!")
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	outputList := []int{}

	inputLen := len(inputList)

	outChan := make(chan int, inputLen)

	var wg sync.WaitGroup

	wg.Add(inputLen)

	for _, i := range inputList {
		go worker(i, outChan, &wg)
	}

	wg.Wait()

	// close channel after worker done, retrieve the result
	// will got deadlock if the channel not closed or use defer to close channel
	close(outChan)
	for oc := range outChan {
		outputList = append(outputList, oc)
	}

	log.Print("[main] out: ", outputList)
	log.Println("[main] bye")
}

func worker(in int, outChan chan<- int, wg *sync.WaitGroup) {
	out := 2 * in
	log.Println("[worker]", in, "=>", out)

	outChan <- out

	defer wg.Done()
}

//fan in -> spawn the worker first than publish message to channel
