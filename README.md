# Soan: the "better C" is just C

**Soan** is a subset of Go that transpiles to regular C with zero runtime.

You write valid Go code — structs, methods, interfaces, slices, multiple returns, defer — and get plain C11 code as output. There's no garbage collector or reference counting. Everything is stack-allocated: slices have a fixed capacity, strings are immutable pointer-length pairs, and interfaces are inline vtable structs.

There are no maps, channels, goroutines, closures, or generics. Instead, you get a language that feels like Go, uses standard Go tools for type-checking, and compiles to C code you could maintain by hand. C interop is first-class: if you declare a function without a body, it's treated as an extern; if you mark type as extern, it comes from your own headers. CGO is not used — Soan provides zero-cost interop with C.

Soan is for people who want Go's syntax and ergonomics for the kind of programs C is good at.
