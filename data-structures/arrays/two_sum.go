package arrays

// twoSum finds two numbers in the given array that add up to the target value.
// It returns the indices of the two numbers as a slice.
// If no such numbers are found, it returns an empty slice.
func twoSum(nums []int, target int) []int {
	visited := make(map[int]int, len(nums))

	for i, number := range nums {
		diff := target - number

		if j, ok := visited[diff]; ok {
			return []int{i, j}
		}

		visited[number] = i
	}

	return []int{}
}
