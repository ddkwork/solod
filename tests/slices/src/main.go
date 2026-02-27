package main

func main() {
	nums := [...]int{1, 2, 3, 4, 5}

	s1 := nums[:]
	s1[1] = 200
	_ = s1

	s2 := nums[2:]
	_ = s2

	s3 := nums[:3]
	_ = s3

	s4 := nums[1:4]
	_ = s4

	n := copy(s4, s1) // n == 3
	_ = n

	strSlice := []string{"a", "b", "c"}
	sLen := len(strSlice) // sLen == 3
	_ = sLen

	twoD := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	x := twoD[0][1] // x == 2
	_ = x
}
