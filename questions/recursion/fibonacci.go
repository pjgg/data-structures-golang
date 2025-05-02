package recursion

// The Fibonacci sequence is de�ned as follows:
// the �rst number of the sequence is 0 , the
// second number is 1 , and the nth number is
// the sum of the (n - 1)th and (n - 2)th
// numbers. Write a function that takes in an
// integer n and returns the nth Fibonacci
// number.

// Important note: the Fibonacci sequence is
// often de�ned with its �rst two numbers as
// F0 = 0 and F1 = 1 . For the purpose of
// this question, the �rst Fibonacci number is
// F0 ; therefore, getNthFib(1) is equal to
// F0 , getNthFib(2) is equal to F1 , etc..

func GetNthFib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return GetNthFib(n-1) + GetNthFib(n-2)
}
