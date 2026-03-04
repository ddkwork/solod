package main

func main() {
	i := 1
	for i <= 3 {
		println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		println(j)
	}

	for k := range 3 {
		println("range", k)
	}

	for {
		println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		println(n)
	}
}
