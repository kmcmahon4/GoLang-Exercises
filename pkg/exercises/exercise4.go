package exercises

import (
	"strconv"
	"strings"
)

// Ex004 takes a string of comma-seperated numbers and returns a slice of int
// "1,2,3,4,5,6,7" -> [1 2 3 4 5 6 7]

func Exercise4(inputString string) []int {
	numbers := strings.Split(inputString, ",")
	length := len(numbers)

	result := make([]int, length)

	for i, value := range numbers {
		result[i], _ = strconv.Atoi(value)
	}

	return result

}
