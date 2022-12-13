package main

import (
	"fmt"
	"os"
	"os/exec"
)

// docker commands
// docker run <image> <cmd> <params>
// docker run main.go   run <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("Help!")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
