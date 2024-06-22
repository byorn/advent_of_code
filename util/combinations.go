package util

func GetCombinations(n int, k int) [][]int {
	var result [][]int
	var comb []int
	return backTrack(1, comb, result, n, k)
}

func backTrack(start int, comb []int, result [][]int, n int, k int) [][]int {
	if len(comb) == k {
		var newVar []int
		newVar = append(newVar, comb...)
		result = append(result, newVar)
		return result
	}

	for i := start; i <= n; i++ {
		comb := append(comb, i)
		result = backTrack(i+1, comb, result, n, k)
		comb = comb[:len(comb)-1]
	}
	return result
}
