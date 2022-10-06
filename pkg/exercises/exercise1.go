package exercises

/*
Exercise 001:
Write a program which will find all such numbers
which are divisible by 7 but are not a multiple of 5,
between 2000 and 3200 (both included).
*/
func Exercise1() []int {
	numbers := make([]int, 0)

	for i := 2000; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			numbers = append(numbers, i)
		}

	}

	return numbers
}
