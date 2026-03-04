package main

func main() {
	var a [5]int
	_ = a

	a[4] = 100
	x := a[4]
	_ = x

	l := len(a)
	_ = l

	b := [5]int{1, 2, 3, 4, 5}
	_ = b

	c := [...]int{1, 2, 3, 4, 5}
	_ = c

	d := [...]int{100, 3: 400, 500}
	_ = d
}
