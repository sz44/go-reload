package main

import (
	"fmt"
	"os"
	"os/exec"
)

// cli tool goreload program.go
// step 1: given program run it
// ../oop/adapter.go
// go run main.go ../oop/adapter.go
// check if server running - or not always stop it
// programFile := "../oop/adapter.go"
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: goreload <program>")
		os.Exit(1)
	}
	programFile := os.Args[1]

	cmd := exec.Command("go", "run", programFile)

	fmt.Println("Running program:", cmd.Path)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("out error: ", err)
	}

	fmt.Printf("%s", out)
}
