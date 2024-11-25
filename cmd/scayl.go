package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var processCount atomic.Uint32

func main() {
	numberOfProcessesAsStr := os.Args[1]
	numberOfProcesses, _ := strconv.Atoi(numberOfProcessesAsStr)

	start := time.Now()

	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < numberOfProcesses; i++ {
		wg.Add(1)
		go launchProcess(ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	totalDuration := time.Since(start) / time.Millisecond

	fmt.Println(fmt.Sprintf("Total duration: %dms", totalDuration))
	fmt.Println(fmt.Sprintf("Number of processes: %d", processCount.Load()))
}

func launchProcess(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	processCount.Add(1)
	currentProcess := processCount.Load()

	ch <- fmt.Sprintf("Process %d starting", currentProcess)
	duration := process()
	ch <- fmt.Sprintf("Process %d done, took %dms", currentProcess, duration)
	/*
		fmt.Println(fmt.Sprintf("Process %d starting", processCount.Add(1)))
		duration := process()
		fmt.Println(fmt.Sprintf("Process %d ended, took %dms", processCount, duration))
		fmt.Println()

		processCount++
	*/
}

func process() int {
	minMs := 50
	maxMs := 100

	randomMs := rand.Intn(maxMs-minMs) + minMs

	time.Sleep(time.Duration(randomMs) * time.Millisecond)

	return randomMs
}
