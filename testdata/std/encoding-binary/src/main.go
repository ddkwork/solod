package main

import "solod.dev/so/encoding/binary"

func main() {
	{
		// Big endian.
		const n1 uint64 = 0x0123456789abcdef
		const n2 uint64 = 0xfedcba9876543210
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, n1)
		if got := binary.BigEndian.Uint64(buf); got != n1 {
			panic("BigEndian: invalid decoded n1")
		}
		buf = binary.BigEndian.AppendUint64(buf[:0], n2)
		if got := binary.BigEndian.Uint64(buf); got != n2 {
			panic("BigEndian: invalid decoded n2")
		}
	}
	{
		// Little endian.
		const n1 uint64 = 0x0123456789abcdef
		const n2 uint64 = 0xfedcba9876543210
		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, n1)
		if got := binary.LittleEndian.Uint64(buf); got != n1 {
			panic("LittleEndian: invalid decoded n1")
		}
		buf = binary.LittleEndian.AppendUint64(buf[:0], n2)
		if got := binary.LittleEndian.Uint64(buf); got != n2 {
			panic("LittleEndian: invalid decoded n2")
		}
	}
}
