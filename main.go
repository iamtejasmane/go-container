package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker commands
// docker run <image> <cmd> <params>
// docker run main.go   run <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Help!")
	}
}

func run() {
	fmt.Printf("Running %v as PID %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	must(cmd.Run())

	syscall.Sethostname([]byte("Pune"))
}

func child() {
	fmt.Printf("Running %v as PID %d\n", os.Args[2:], os.Getpid())
	syscall.Sethostname([]byte("Pune"))

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
