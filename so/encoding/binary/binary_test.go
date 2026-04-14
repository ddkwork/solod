// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binary

import (
	"testing"

	"solod.dev/so/math"
)

func TestByteOrder(t *testing.T) {
	type byteOrder interface {
		ByteOrder
		AppendByteOrder
	}
	buf := make([]byte, 8)
	for _, order := range []byteOrder{LittleEndian, BigEndian} {
		const offset = 3
		for _, value := range []uint64{
			0x0000000000000000,
			0x0123456789abcdef,
			0xfedcba9876543210,
			0xffffffffffffffff,
			0xaaaaaaaaaaaaaaaa,
			math.Float64bits(math.Pi),
			math.Float64bits(math.E),
		} {
			want16 := uint16(value)
			order.PutUint16(buf[:2], want16)
			if got := order.Uint16(buf[:2]); got != want16 {
				t.Errorf("PutUint16: Uint16 = %v, want %v", got, want16)
			}
			buf = order.AppendUint16(buf[:offset], want16)
			if got := order.Uint16(buf[offset:]); got != want16 {
				t.Errorf("AppendUint16: Uint16 = %v, want %v", got, want16)
			}
			if len(buf) != offset+2 {
				t.Errorf("AppendUint16: len(buf) = %d, want %d", len(buf), offset+2)
			}

			want32 := uint32(value)
			order.PutUint32(buf[:4], want32)
			if got := order.Uint32(buf[:4]); got != want32 {
				t.Errorf("PutUint32: Uint32 = %v, want %v", got, want32)
			}
			buf = order.AppendUint32(buf[:offset], want32)
			if got := order.Uint32(buf[offset:]); got != want32 {
				t.Errorf("AppendUint32: Uint32 = %v, want %v", got, want32)
			}
			if len(buf) != offset+4 {
				t.Errorf("AppendUint32: len(buf) = %d, want %d", len(buf), offset+4)
			}

			want64 := uint64(value)
			order.PutUint64(buf[:8], want64)
			if got := order.Uint64(buf[:8]); got != want64 {
				t.Errorf("PutUint64: Uint64 = %v, want %v", got, want64)
			}
			buf = order.AppendUint64(buf[:offset], want64)
			if got := order.Uint64(buf[offset:]); got != want64 {
				t.Errorf("AppendUint64: Uint64 = %v, want %v", got, want64)
			}
			if len(buf) != offset+8 {
				t.Errorf("AppendUint64: len(buf) = %d, want %d", len(buf), offset+8)
			}
		}
	}
}
