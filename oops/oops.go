package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/samber/oops"
)

func a() error {
	err := b()
	if err != nil {
		return err
	}
	return nil
}

func b() error {
	err := f("foo.txt")
	if err != nil {
		return err
	}
	return nil
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
	err := f("bar.txt")
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// oops.SourceFragmentsHidden = false

	fmt.Printf("a: %+v\n", a())

	err := a()
	slog.Error(err.Error(), slog.Any("error", err))
}
