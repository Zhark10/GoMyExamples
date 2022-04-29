package channel

import (
	"fmt"
	"sync"
)

var counter int = 0

func RunProcessWithMutex() {
	ch := make(chan bool)
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		go work(i, ch, &mutex)
		<-ch
	}
}

func work(i int, ch chan bool, mt *sync.Mutex) {
	mt.Lock()
	counter = 0
	for index := 0; index < 3; index++ {
		counter++
		fmt.Println("Goroutine", i, "-", counter)
	}
	mt.Unlock()
	ch <- true
}
