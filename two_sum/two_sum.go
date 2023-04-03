package two_sum

// twoSum1 暴力解法
func twoSum1(nums []int, target int) []int {
	for i, v := range nums {
		for k := i + 1; k < len(nums); k++ {
			if target-v == nums[k] {
				return []int{i, k}
			}
		}
	}
	return []int{}
}

// twoSum 两数之和
func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	m := make(map[int]int)
	for i, k := range nums {
		if value, exist := m[target-k]; exist {
			result = append(result, value)
			result = append(result, i) // 当前的索引
		}
		m[k] = i
	}
	return result
}
