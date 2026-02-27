package main

func main() {
	if 7%2 == 0 {
		println("7 is even")
	} else {
		println("7 is odd")
	}

	if 8%2 == 0 || 7%2 == 0 {
		println("either 8 or 7 are even")
	}

	if 1 == 2-1 && (2 == 1+1 || 3 == 6/2) && !(4 != 2*2) {
		println("all conditions are true")
	}

	if 9%3 == 0 {
		println("9 is divisible by 3")
	} else if 9%2 == 0 {
		println("9 is divisible by 2")
	} else {
		println("9 is not divisible by 3 or 2")
	}

	if num := 9; num < 0 {
		println(num, "is negative")
	} else if num < 10 {
		println(num, "has 1 digit")
	} else {
		println(num, "has multiple digits")
	}
}
