package arrays

// Write a function that takes in a non-empty
// array of integers that are sorted in ascending
// order and returns a new array of the same
// length with the squares of the original
// integers also sorted in ascending orde

// array = [1, 2, 3, 5, 6, 8, 9]
// [1, 4, 9, 25, 36, 64, 81]

func SortedSquaredArray(array []int) []int {
	output := make([]int, len(array))
	end := len(array) - 1
	index := len(array) - 1
	start := 0

	for end >= start {
		startElem := array[start]
		endElem := array[end]

		if startElem < 0 {
			startElem = startElem * -1
		}

		if endElem < 0 {
			endElem = endElem * -1
		}

		if startElem <= endElem {
			output[index] = endElem * endElem
			index--
			end--
		} else {
			output[index] = startElem * startElem
			start++
			index--
		}
	}

	return output
}
