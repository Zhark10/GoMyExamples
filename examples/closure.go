package examples

import "fmt"

func Closure() {
	newCounter := counter(1)
	fmt.Println(newCounter(1))
	fmt.Println(newCounter(1))
	fmt.Println(newCounter(1))
	fmt.Println(newCounter(20))
}

func counter(initial_value int) func(int) int {
	count := initial_value

	return func(inc_value int) int {
		count += inc_value
		return count
	}
}
