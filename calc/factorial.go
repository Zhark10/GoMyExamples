package calc

import "fmt"

func Factorial(n int) {
	if(n < 1){
		fmt.Println("Unvalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++{
			result *= i
	}
	fmt.Println(n, "-", result)
}
	