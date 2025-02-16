package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// cli tool goreload program.go
// step 1: given program run it
// ../oop/adapter.go
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: goreload <program>")
		os.Exit(1)
	}
	programFile := os.Args[1]

	cmd := exec.Command(programFile)

	fmt.Println("Running program:", cmd.Path)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	var b strings.Builder
	b.Write(out)

	fmt.Println(b.String())
}
