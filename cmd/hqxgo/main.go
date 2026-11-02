// Command hqxgo enlarges an image four times with the HQ4x filter.
package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"

	hqx "github.com/donatj/hqxgo"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "hqxgo:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("usage: hqxgo input output.png")
	}
	input, err := os.Open(args[0])
	if err != nil {
		return fmt.Errorf("open input: %w", err)
	}
	src, _, err := image.Decode(input)
	closeErr := input.Close()
	if err != nil {
		return fmt.Errorf("decode input: %w", err)
	}
	if closeErr != nil {
		return fmt.Errorf("close input: %w", closeErr)
	}

	dst := hqx.Scale4x(src)
	output, err := os.Create(args[1])
	if err != nil {
		return fmt.Errorf("open output: %w", err)
	}
	err = png.Encode(output, dst)
	closeErr = output.Close()
	if err != nil {
		return fmt.Errorf("encode output: %w", err)
	}
	if closeErr != nil {
		return fmt.Errorf("close output: %w", closeErr)
	}
	return nil
}
