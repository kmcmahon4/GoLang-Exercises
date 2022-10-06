package exercises

/*
Write a program which can compute the factorial of a given number.
Suppose the following input is supplied to the program: 8
Then, the output should be: 40320
*/

func Exercise2(number int) (result uint64) {
	result = 1
	for i := 1; i <= number; i++ {
		result *= uint64(i)
	}

	return result

}
