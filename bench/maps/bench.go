package main

import (
	"solod.dev/so/fmt"
	"solod.dev/so/maps"
	"solod.dev/so/mem"
	"solod.dev/so/strings"
	"solod.dev/so/testing"
)

//so:embed bench.h
var bench_h string

var arena *mem.Arena

//so:extern nodecay
var (
	sinkInt  int
	sinkBool bool
)

// nKeys is the number of map keys to use in benchmarks.
const nKeys = 1024

// strKeys holds pre-generated string keys for string benchmarks.
var strKeys []string

func initStrKeys() {
	strKeys = mem.AllocSlice[string](nil, nKeys, nKeys)
	buf := fmt.NewBuffer(32)
	for i := range nKeys {
		strKeys[i] = strings.Clone(nil, fmt.Sprintf(buf, "key-%d", i))
	}
}

func freeStrKeys() {
	for i := range nKeys {
		mem.FreeString(nil, strKeys[i])
	}
	mem.FreeSlice(nil, strKeys)
}

func IntSet(b *testing.B) {
	a := b.Allocator()
	for b.Loop() {
		m := maps.New[int, int](a, 0)
		for i := range nKeys {
			m.Set(i, i)
		}
		m.Free()
		if arena != nil {
			arena.Reset()
		}
	}
}

func IntGet(b *testing.B) {
	m := maps.New[int, int](nil, nKeys)
	for i := range nKeys {
		m.Set(i, i)
	}
	defer m.Free()
	for b.Loop() {
		for i := range nKeys {
			sinkInt = m.Get(i)
		}
	}
}

func IntHas(b *testing.B) {
	m := maps.New[int, int](nil, nKeys)
	for i := range nKeys {
		m.Set(i, i)
	}
	defer m.Free()
	for b.Loop() {
		for i := range nKeys {
			sinkBool = m.Has(i)
		}
	}
}

func IntDelete(b *testing.B) {
	a := b.Allocator()
	for b.Loop() {
		m := maps.New[int, int](a, nKeys)
		for i := range nKeys {
			m.Set(i, i)
		}
		for i := range nKeys {
			m.Delete(i)
		}
		m.Free()
		if arena != nil {
			arena.Reset()
		}
	}
}

func StrSet(b *testing.B) {
	a := b.Allocator()
	for b.Loop() {
		m := maps.New[string, int](a, 0)
		for i := range nKeys {
			m.Set(strKeys[i], i)
		}
		m.Free()
		if arena != nil {
			arena.Reset()
		}
	}
}

func StrGet(b *testing.B) {
	m := maps.New[string, int](nil, nKeys)
	for i := range nKeys {
		m.Set(strKeys[i], i)
	}
	defer m.Free()
	for b.Loop() {
		for i := range nKeys {
			sinkInt = m.Get(strKeys[i])
		}
	}
}

func StrHas(b *testing.B) {
	m := maps.New[string, int](nil, nKeys)
	for i := range nKeys {
		m.Set(strKeys[i], i)
	}
	defer m.Free()
	for b.Loop() {
		for i := range nKeys {
			sinkBool = m.Has(strKeys[i])
		}
	}
}

func StrDelete(b *testing.B) {
	a := b.Allocator()
	for b.Loop() {
		m := maps.New[string, int](a, nKeys)
		for i := range nKeys {
			m.Set(strKeys[i], i)
		}
		for i := range nKeys {
			m.Delete(strKeys[i])
		}
		m.Free()
		if arena != nil {
			arena.Reset()
		}
	}
}

func StackSet(b *testing.B) {
	nKeys := 128
	for b.Loop() {
		stackSet(nKeys) // alloca only frees when the function returns
	}
}

func stackSet(nKeys int) {
	m := make(map[string]int, nKeys)
	for i := range nKeys {
		m[strKeys[i]] = i
	}
	sinkInt = m[strKeys[0]]
}

func StackGet(b *testing.B) {
	nKeys := 128
	m := make(map[string]int, nKeys)
	for i := range nKeys {
		m[strKeys[i]] = i
	}
	for b.Loop() {
		for i := range nKeys {
			sinkInt = m[strKeys[i]]
		}
	}
}

func StackHas(b *testing.B) {
	nKeys := 128
	m := make(map[string]int, nKeys)
	for i := range nKeys {
		m[strKeys[i]] = i
	}
	for b.Loop() {
		for i := range nKeys {
			_, sinkBool = m[strKeys[i]]
		}
	}
}

func main() {
	initStrKeys()
	defer freeStrKeys()

	benchs := []testing.Benchmark{
		{Name: "IntSet", F: IntSet},
		{Name: "IntGet", F: IntGet},
		{Name: "IntHas", F: IntHas},
		{Name: "IntDelete", F: IntDelete},
		{Name: "StrSet", F: StrSet},
		{Name: "StrGet", F: StrGet},
		{Name: "StrHas", F: StrHas},
		{Name: "StrDelete", F: StrDelete},
	}

	fmt.Println("Malloc-based allocator:")
	testing.RunBenchmarks(mem.System, benchs)

	fmt.Println("Arena allocator:")
	const size = 4 << 20
	buf := mem.AllocSlice[byte](nil, size, size)
	defer mem.FreeSlice(nil, buf)
	a := mem.NewArena(buf[:])
	arena = &a
	testing.RunBenchmarks(arena, benchs)

	stackBenchs := []testing.Benchmark{
		{Name: "StackSet", F: StackSet},
		{Name: "StackGet", F: StackGet},
		{Name: "StackHas", F: StackHas},
	}
	fmt.Println("Stack-based map:")
	testing.RunBenchmarks(mem.System, stackBenchs)
}
