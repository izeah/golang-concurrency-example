package main

import (
	"fmt"
	"golang-concurrency/channel"
	"golang-concurrency/goroutine"
	"golang-concurrency/waitgroup"
)

func main() {
	fmt.Println("==================================================================")
	fmt.Println("			GOLANG CONCURRENCY			  ")
	fmt.Println("	Goroutine, Channel, Mutex, WaitGroup, WorkerPool")
	fmt.Println("==================================================================")

	// Goroutine
	fmt.Println()
	fmt.Println("// Goroutine")
	goroutine.Run()

	// Channel
	fmt.Println()
	fmt.Println("// Channel")
	channel.Run()

	// WaitGroup
	fmt.Println()
	fmt.Println("// WaitGroup")
	waitgroup.Run()
}
