package arrays

// Given two non-empty arrays of integers,
// write a function that determines whether the
// second array is a subsequence of the �rst
// one.
// A subsequence of an array is a set of
// numbers that aren't necessarily adjacent in
// the array but that are in the same order as
// they appear in the array. For instance, the
// numbers [1, 3, 4] form a subsequence
// f the array [1, 2, 3, 4] , and so do the
// numbers [2, 4] . Note that a single
// number in an array and the array itself are
// both valid subsequences of the array

func IsValidSubsequence(array []int, sequence []int) bool {
	i := 0
	for _, val := range array {
		if val == sequence[i] {
			i++
		}

		if len(sequence) == i {
			return true
		}
	}

	return false
}
