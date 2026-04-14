package main

import "solod.dev/so/math/rand"

func main() {
	{
		// Int.
		pcg := rand.NewPCG(1, 2)
		r := rand.New(&pcg)
		n1 := r.Int()
		if n1 < 0 {
			panic("negative Int()")
		}
		n2 := r.Int()
		if n2 < 0 {
			panic("negative Int()")
		}
		if n1 == n2 {
			panic("same Int() twice in a row")
		}
		println(n1, n2)
	}
	{
		// Float64.
		pcg := rand.NewPCG(1, 2)
		r := rand.New(&pcg)
		f1 := r.Float64()
		if f1 < 0 || f1 >= 1 {
			panic("Float64() out of range")
		}
		f2 := r.Float64()
		if f2 < 0 || f2 >= 1 {
			panic("Float64() out of range")
		}
		if f1 == f2 {
			panic("same Float64() twice in a row")
		}
		println(f1, f2)
	}
	{
		// Global functions.
		n1 := rand.IntN(100)
		if n1 < 0 || n1 >= 100 {
			panic("IntN() out of range")
		}
		n2 := rand.IntN(100)
		if n2 < 0 || n2 >= 100 {
			panic("IntN() out of range")
		}
		println(n1, n2)
	}
}
