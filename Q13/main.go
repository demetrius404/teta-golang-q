package main

import (
	"fmt"
	"sync"
)

func printWorker(i int) {
	fmt.Print(i)
}

func printWorkerChannel(doneCh chan bool, v int) {
	fmt.Print(v)
	doneCh <- true
}

func MutexQ12() {
	var mutex sync.Mutex
	fmt.Print("Q13 (with Mutex): ")
	
	for i := 0; i < 11; i++ {
		mutex.Lock()
		go func(i int) {
			defer mutex.Unlock()
			printWorker(i)
		}(i)
	}
	fmt.Print("\n")
}

func ChannelQ12() {
	// channel for synchronization
	fmt.Print("Q13 (with Channel): ")
	blockCh := make(chan bool)
	for i := 0; i < 10; i++ {
		go printWorkerChannel(blockCh, i)
		<- blockCh
	}
	
	close(blockCh)

	fmt.Print("\n")
}


func main() {
	fmt.Println("--- Q13 ---")
	MutexQ12()
	ChannelQ12()
}

// Q13 (default) output: (empty or random partial output)
// Q13 (with WaitGroup) output: 9012345678 (depend from runtime)
// Q13 (with Mutex) output: 01213456789
// Q13 (with Channel) output: 01213456789 (channel for synchronization)