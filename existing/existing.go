package main

import (
	"fmt"
	"os"
)

func a() error {
	err := b()
	if err != nil {
		return fmt.Errorf("b failed: %w", err)
	}
	return nil
}

func b() error {
	err := f("foo.txt")
	if err != nil {
		return fmt.Errorf("c failed: %w", err)
	}
	return nil
}

func f(file string) error {
	f, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	fmt.Printf("read: %s\n", string(f))
	return nil
}

func c() error {
	err := f("bar.txt")
	if err != nil {
		return fmt.Errorf("d failed: %w", err)
	}
	return nil
}

func main() {
	fmt.Printf("a: %s\n", a())
	fmt.Printf("b: %s\n", b())
	fmt.Printf("c: %s\n", c())
}
