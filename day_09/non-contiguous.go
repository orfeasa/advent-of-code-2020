// I initially did't pay attention to the word 'contiguous' in the problem description
// so I had to solve a much harder problem (and is too slow to solve a problem of the size of the input)
// nevertheless, here is what I did :p

// countSubsetSums returns the amount of subsets of numbers[start:] that sum exactly to target
func countSubsetSums(numbers []int, start, target int, memo map[[2]int]int) int {

	if start >= len(numbers) {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}

	if _, contains := memo[[2]int{start, target}]; !contains {
		count := countSubsetSums(numbers, start+1, target, memo)
		count += countSubsetSums(numbers, start+1, target-numbers[start], memo)
		memo[[2]int{start, target}] = count
	}
	return memo[[2]int{start, target}]
}

// subsetSum return a subset of numbers that is equal to target
func subsetSumMemo(numbers []int, target int, memo map[[2]int]int) []int {
	var subset []int

	for ind, val := range numbers {
		// check if there's still a solution if we include numbers[ind]
		if countSubsetSums(numbers, ind+1, target-val, memo) > 0 {
			subset = append(subset, val)
			target -= val
		}
	}
	return subset
}

// subsetSum return a subset of numbers that is equal to target
func subsetSum(numbers []int, target int) []int {
	memo := make(map[[2]int]int)
	return subsetSumMemo(numbers, target, memo)
}
