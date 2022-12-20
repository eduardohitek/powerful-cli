package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func run(proj string, out io.Writer) error {
	if proj == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}
	args := []string{"build", ".", "errors"}
	cmd := exec.Command("go", args...)
	cmd.Dir = proj

	err := cmd.Run()
	if err != nil {
		return &stepErr{step: "go build", msg: "go build failed", cause: err}
	}

	_, err = fmt.Fprintln(out, "Go Build: SUCCESS")
	return err
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
