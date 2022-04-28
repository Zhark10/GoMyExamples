package examples

import "fmt"

func Closure() {
	newCounter := counter(1)
	fmt.Println(newCounter(1))
	fmt.Println(newCounter(1))
	fmt.Println(newCounter(1))
	fmt.Println(newCounter(20))
}

func counter(initialValue int) func(int) int {
	count := initialValue

	return func(incValue int) int {
		count += incValue
		return count
	}
}
