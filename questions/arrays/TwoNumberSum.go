package arrays

// Write a function that takes in a non-empty
// array of distinct integers and an integer
// representing a target sum. If any two
// numbers in the input array sum up to the
// target sum, the function should return them
// in an array, in any order. If no two numbers
// sum up to the target sum, the function
// should return an empty array.
// Note that the target sum has to be obtained
// by summing two di�erent integers in the
// array; you can't add a single integer to itself
// in order to obtain the target sum.
// You can assume that there will be at most
// one pair of numbers summing up to the
// target sum.

func TwoNumberSum(array []int, target int) []int {
	storedValues := make(map[int]int)
	for _, currentNum := range array {
		storedValues[currentNum] = currentNum
	}

	result := []int{}
	for _, value := range storedValues {
		numNeeded := target - value
		if value == numNeeded {
			continue
		}

		if _, exists := storedValues[numNeeded]; exists {
			result = append(result, value)
			result = append(result, numNeeded)
			break
		}
	}

	return result
}
