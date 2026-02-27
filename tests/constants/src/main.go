package main

const s string = "constant"

func sin(x float64) float64 {
	return x // stub
}

func main() {
	println(s)

	const n = 500000000

	const d = 3e20 / n
	println(d)

	println(int64(d))

	println(sin(n))
}
