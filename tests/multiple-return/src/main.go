package main

func vals() (int, int) {
	return 3, 7
}

func swap(x int, y int) (int, int) {
	return y, x
}

func divide(x int, y int) (res int, mod int) {
	return x / y, x % y
}

func main() {
	a, b := vals()
	b, a = swap(a, b)
	_ = a
	_ = b

	d, m := divide(7, 3)
	_ = d
	_ = m

	_, c := vals()
	_ = c
}
