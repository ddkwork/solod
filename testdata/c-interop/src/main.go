package main

import (
	"github.com/nalgeon/solod/so/c"
	"github.com/nalgeon/solod/so/c/stdlib"
)

func main() {
	{
		// c.String: convert C string to So string.
		ptr := stdlib.Getenv("PATH")
		path := c.String(ptr)
		if len(path) == 0 {
			panic("want non-empty PATH")
		}
	}
	{
		// c.String: nil pointer returns empty string.
		ptr := stdlib.Getenv("SOLOD_NONEXISTENT_VAR")
		s := c.String(ptr)
		if len(s) != 0 {
			panic("want empty string for nil")
		}
	}
	{
		// c.Bytes: wrap a raw buffer into []byte.
		buf := stdlib.Malloc(4)
		if buf == nil {
			panic("malloc failed")
		}
		ptr := any(buf).(*byte)
		*ptr = 'H'
		slice := c.Bytes(ptr, 4)
		if len(slice) != 4 {
			panic("want len == 4")
		}
		if slice[0] != 'H' {
			panic("want slice[0] == 'H'")
		}
		stdlib.Free(buf)
	}
}
