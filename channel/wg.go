package channel

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroupExample() {
	var wg sync.WaitGroup
	const count int = 10

	waitGroup(&wg, count)

	wg.Wait()
}

func goroutine(wg *sync.WaitGroup, index int) {
	defer wg.Done()
	fmt.Println("Goroutine", index, "выполняется")
	time.Sleep(time.Second)
	fmt.Println("Goroutine", index, "завершена")
}

func waitGroup(wg *sync.WaitGroup, num int) {
	for i := 0; i < num; i++ {
		wg.Add(1)
		go goroutine(wg, i)
	}
}