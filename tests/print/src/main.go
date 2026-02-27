package main

type person struct {
	name string
}

func main() {
	var a int = 42
	var b float64 = 3.14
	var c bool = true
	var d byte = 'x'
	var e string = "hello"
	alice := person{name: "alice"}
	var f = &alice
	println(a, b, c, d, e, f)

	// Complex types are not supported.
	// arr := [3]int{1, 2, 3}
	// println(arr)
	// alice := person{name: "alice"}
	// println(alice)
}
