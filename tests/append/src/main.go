package main

func main() {
	nums := make([]int, 3)
	n := nums[1] // 0
	nums[1] = 42
	l1 := len(nums) // 3
	c1 := cap(nums)

	nums = make([]int, 0, 3)
	nums = append(nums, 1)
	nums = append(nums, 2, 3)
	l2 := len(nums) // 3
	c2 := cap(nums) // 3

	// Resizing slices beyond their initial capacity with append() panics.
	// nums = append(nums, 4)

	_ = n
	_ = l1
	_ = c1
	_ = l2
	_ = c2
}
