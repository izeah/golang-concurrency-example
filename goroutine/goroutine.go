package goroutine

import (
	"fmt"
	"runtime"
)

func printGoroutine(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func Run() {
	runtime.GOMAXPROCS(2)

	go printGoroutine(5, "Foo")
	printGoroutine(5, "Lorem Ipsum")

	var input string
	fmt.Scanln(&input)
}
