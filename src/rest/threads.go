package main

import (
	"fmt"
	"time"
)

func printMessage(){
	fmt.Println("Hello from Goroutine")
}

func main1(){
	go printMessage();
	fmt.Println("Hello from main function")
	time.Sleep(time.Second)
}