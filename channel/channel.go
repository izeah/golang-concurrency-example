package channel

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}

	ch <- max
}

func Run() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("Hello, %s", who)
		messages <- data // send data via channel
	}

	go sayHelloTo("lorem")
	go sayHelloTo("lorem ipsum")
	go sayHelloTo("lorem ipsum dolor sit amet")

	var message1 = <-messages // receiving data from channel
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	fmt.Println()

	// Channel as parameter data type
	for _, each := range []string{"foo", "bar", "baz"} {
		go func(who string) {
			var data = fmt.Sprintf("Hello, %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}

	fmt.Println()

	// Buffered Channel
	messagesBuffer := make(chan int, 2)

	go func() {
		for {
			i := <-messagesBuffer
			fmt.Println("receive data :", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data :", i)
		messagesBuffer <- i
	}

	// Channel Select
	var numbers = []int{3, 6, 1, 0, 2, 7, 6, 9, 4, 5, 8}
	fmt.Println("numbers :", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg\t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Avg\t: %d \n", max)
		}
	}
}
