func moveZeroes(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i] == 0 {
			k := i
			target := nums[k]
			for k < j {
				nums[k] = nums[k+1]
				k++
			}
			nums[j] = target
			j--
			continue
		}
		i++
	}
}