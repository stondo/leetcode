package twosums

func twoSum(nums []int, target int) []int {
	numsLen := len(nums)

	for i := 0; i < numsLen; i++ {
		for k := numsLen - 1; k >= 0; k-- {
			if i != k && nums[i]+nums[k] == target {
				return []int{i, k}
			}
		}
	}

	return nil
}
