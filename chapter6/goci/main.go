package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func run(proj string, out io.Writer) error {
	if proj == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}
	pipeline := make([]step, 1)
	pipeline[0] = newStep(
		"go build",
		"go",
		"Go Build: SUCCESS",
		proj,
		[]string{"build", ".", "errors"},
	)

	for _, p := range pipeline {
		msg, err := p.execute()
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(out, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	proj := flag.String("p", "", "Project directory")
	flag.Parse()

	err := run(*proj, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
