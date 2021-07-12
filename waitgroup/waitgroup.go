package waitgroup

import (
	"fmt"
	"runtime"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {
	defer wg.Done()
	fmt.Println(message)
}

func Run() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		var data = fmt.Sprintf("data %d", i)

		wg.Add(1)
		go doPrint(&wg, data)
	}

	wg.Wait()
	fmt.Println("waitgroup selesai")
}
