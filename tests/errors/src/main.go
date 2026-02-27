package main

import (
	"errors"
)

var ErrOutOfTea = errors.New("no more tea available")

func makeTea(arg int) error {
	if arg == 42 {
		return ErrOutOfTea
	}
	return nil
}

func main() {
	err := makeTea(7)
	if err == nil {
		println("err == nil")
	}

	err = makeTea(42)
	if err != nil {
		println("err != nil")
	}
	if err == ErrOutOfTea {
		println("err == ErrOutOfTea")
	}

	// Not supported: errors can only be defined at package level.
	// errNotSupported := errors.New("operation not supported")

	// Dynamic errors are also not supported.
	// errNotSupported := fmt.Errorf("not supported: %d", 42)
}
