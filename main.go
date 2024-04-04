package main

import (
	"fmt"
	"sync"
	"time"
)

func qd() {
	time.Sleep(5 * time.Second)
	fmt.Print("Quan dao")
}

type APIData struct {
	Data string
}

func FetchAPI1(gr *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("Task 1 done")
	gr.Done()
}

func FetchAPI2(gr *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Task 2 done")
	gr.Done()
}

func FetchAPI3(gr *sync.WaitGroup) {
	time.Sleep(4 * time.Second)
	fmt.Println("Task 3 done")
	gr.Done()
}

func main() {
	//goroutine
	// -----------------------------------------------------main--------->
	//  |----------------------------------------------------F1------------>
	//  |-----------------------------------------------------F2----------->
	// Main <- Channel <- Routine  || bufferedChannel
	groupFetch := sync.WaitGroup{}
	groupFetch.Add(3)
	go FetchAPI1(&groupFetch)
	go FetchAPI2(&groupFetch)
	go FetchAPI3(&groupFetch)
	groupFetch.Wait()
	fmt.Println("All tasks done")
}
