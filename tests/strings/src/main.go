package main

func main() {
	str := "Hi 世界!"

	// Loop over bytes.
	for i := 0; i < len(str); i++ {
		chr := str[i]
		println("i =", i, "chr =", chr)
	}

	// Loop over runes.
	for i, r := range str {
		println("i =", i, "r =", r)
	}
	for i := range str {
		println("i =", i)
	}
	for _, r := range str {
		println("r =", r)
	}

	s1 := "hello"
	s2 := "world"
	if s1 == s2 || s1 == "hello" {
		println("ok")
	}
}
