package channel

import "fmt"

func RunProcesses() {
	channel := make(chan string)

	go gorutine(channel)

	for {
		message, opened := <-channel
		if !opened {
			break
		}
		fmt.Println(message)
	}
}

func gorutine(channel chan string) {
	defer close(channel)
	array := []string{"hello", "World", "sdsdf", "s23234"}
	for _, val := range array {
		func(value string) {
			channel <- value
		}(val)
	}
}
