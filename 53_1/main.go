package _53_1

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func search(nums []int, target int) int {
	l := binarySearch(nums, target)
	if l >= len(nums) || nums[l] != target {
		return 0
	}
	r := binarySearch(nums, target+1)
	return r - l
}
