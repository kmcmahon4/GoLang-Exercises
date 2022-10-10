package main

import (
	"fmt"
	"main/pkg/exercises"
	"main/pkg/leetcode"
)

// Just to drive the execution of the functions
func main() {
	// exercises.Exercise1()
	// exercises.Exercise2(6)
	// exercises.Exercise3(5)
	// exercises.Exercise4("String")
	// exercises.Exercise5()
	exercises.Exercise6(4)
	//leetcode.TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(leetcode.ContainsDuplicate([]int{1, 2, 3, 4}))
}
