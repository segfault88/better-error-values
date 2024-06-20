package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/samber/oops"
)

func a() error {
	return oops.Wrap(b())
}

func b() error {
	return oops.Wrap(f("foo.txt"))
}

func f(file string) error {
	f, err := os.ReadFile(file)
	if err != nil {
		return oops.
			In("f reader").
			Wrap(err)
	}

	fmt.Printf("read: %s\n", string(f))
	return nil
}

func c() error {
	return oops.Wrap(f("bar.txt"))
}

func main() {
	// oops.SourceFragmentsHidden = false

	fmt.Printf("a: %+v\n\n", a())
	fmt.Printf("c: %+v\n\n", c())

	err := a()
	slog.Error(err.Error(), slog.Any("error", err))

	fmt.Printf("\n\n\n")

	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("bad thing was file not found\n")
	} else {
		fmt.Printf("bad thing was some other thing\n")
	}
}
