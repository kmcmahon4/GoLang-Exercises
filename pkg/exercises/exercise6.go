package exercises

import "fmt"

// Calculate the factorial of a given number, and then do it with goroutines and channels

func Exercise6(value int) {
	result := 1

	c := make(chan int)

	go func(value int) {
		for i := value; i > 0; i-- {
			result *= i
		}
		c <- result
		close(c)
	}(value)

	fmt.Printf("Result %d ", <-c)

}
