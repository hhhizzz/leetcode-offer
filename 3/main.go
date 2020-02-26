package _1

//置换法 类似贪心的算法，针对每个位置找到对应的数字，如果中间找到了重复的就返回，特点是空间复杂度为O(1)
func findRepeatNumber(nums []int) int {
    for i := 0; i < len(nums); i++ {
        for nums[i] != i {
            next := nums[i]
            if next == nums[next] {
                return next
            }
            temp := nums[i]
            nums[i] = nums[next]
            nums[next] = temp
        }
    }
    return -1
}
