package main

func main() {
	{
		// Integer arithmetics.
		var a, b, c int = 11, 22, 33
		d := b/a + (a-c)*a + c%b
		d += 10
		d -= 10
		d *= 10
		d /= 2
		d %= 5
		d++
		d--
		_ = d
	}

	{
		// Floating-point arithmetics.
		var x, y, z float64 = 1.1, 2.2, 3.3
		f := x/y + (y-z)*x
		f += 1.0
		f -= 1.0
		f *= 2.0
		f /= 2.0
		f++
		f--
		_ = f
	}

	{
		// Bitwise operations.
		var b1, b2 = 0b1010, 0b1100
		b3 := (b1|b2)&(b1&b2) | (b1 ^ b2)
		b3 = b3 << 2
		b3 = b3 >> 1
		b3 = b3 &^ b1
		_ = b3
		b4 := 0b1010
		b4 |= 0b1100
		b4 &= 0b1100
		b4 ^= 0b1100
		// b4 &^= 0b1010 // not supported
		_ = b4
	}

	{
		// Logical operations and comparisons.
		var a, b, c bool = true, false, true
		d := (a && b) || (b || c) && !a
		_ = d

		x, y, z := 10, 20, 30
		e1 := (x < y) && (y > z) || (x == z)
		_ = e1
		e2 := (x <= y) && (y >= z) || (x != z)
		_ = e2
	}
}
