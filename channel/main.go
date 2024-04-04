package main

import (
	"fmt"
	"time"
)

type APIData struct {
	Data string
}

func FetchAPI1(c chan<- APIData) {
	time.Sleep(3 * time.Second)
	c <- APIData{
		Data: "Data 1",
	}
}

func FetchAPI2(c chan<- APIData) {
	time.Sleep(2 * time.Second)
	c <- APIData{
		Data: "Data 2",
	}
}

func Fetch3() {
	time.Sleep(1 * time.Second)
	fmt.Println("Data 3")
}

func main() {
	//goroutine
	// -----------------------------------------------------main--------->
	//  |----------------------------------------------------F1------------>
	//  |-----------------------------------------------------F2----------->
	// Main <- Channel <- Routine  || bufferedChannel
	c := make(chan APIData, 2)

	go FetchAPI1(c)
	go FetchAPI2(c)

	fmt.Println(<-c)
	Fetch3()
	fmt.Println(<-c)
}
