package main

import (
	"log"
	"sync"
)

func main() {
	fanIn()
	fanOut()
}

func fanOut() {
	log.Println("[fanOut] hi!")
	// inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println("[fanOut] bye")
}

func fanIn() {
	log.Println("[fanIn] hi!")
	inputList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	outputList := []int{}

	inputLen := len(inputList)

	outChan := make(chan int, inputLen)

	var wg sync.WaitGroup

	wg.Add(inputLen)

	for _, i := range inputList {
		go multiplyByTwo(i, outChan, &wg)
	}

	wg.Wait()

	// close channel after worker done, retrieve the result
	// will got deadlock if the channel not closed or use defer to close channel
	close(outChan)
	for oc := range outChan {
		outputList = append(outputList, oc)
	}

	log.Print("[fanIn] out: ", outputList)
	log.Println("[fanIn] bye")
}

/** worker: multiply by 2 **/
func multiplyByTwo(in int, outChan chan<- int, wg *sync.WaitGroup) {
	out := 2 * in
	log.Println("[multiplyByTwo]", in, "=>", out)

	outChan <- out

	defer wg.Done()
}

//fan in -> spawn the worker first than publish message to channel
