// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand_test

import (
	"testing"

	. "solod.dev/so/math/rand"
)

func TestFloat32(t *testing.T) {
	num := int(10e4)
	pcg := NewPCG(1, 2)
	r := New(&pcg)
	for ct := range num {
		f := r.Float32()
		if f >= 1 {
			t.Fatal("Float32() should be in range [0,1). ct:", ct, "f:", f)
		}
	}
}
