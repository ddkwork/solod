package main

import "solod.dev/so/c"

//so:include <math.h>

func main() {
	{
		// Typed C expression.
		nan := c.Val[float64]("NAN")
		if nan == nan {
			panic("nan == nan")
		}
		x := c.Val[float64]("sqrt(49)")
		if x != 7 {
			panic("x != 7")
		}
	}
	{
		// Raw C block.
		var b int
		c.Raw(`
		int a = 7;
		b = a * a;
		b = sqrt(b);
		`)
		if b != 7 {
			panic("b != 7")
		}
	}
}
