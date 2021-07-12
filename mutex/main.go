package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	val int
}

func (c *counter) Add() {
	c.val++
}

func (c *counter) Value() (x int) {
	return c.val
}

func main() {
	fmt.Println("GOLANG CONCURRENCY - MUTEX")
	fmt.Println("------------------------------------------------")
	fmt.Println()

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	var meter counter
	var mtx sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Before using mutex")
	fmt.Println(meter.Value())

	meter.val = 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				mtx.Lock()
				meter.Add()
				mtx.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("After using mutex")
	fmt.Println(meter.Value())
}
