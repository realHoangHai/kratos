package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func stripSlice(slice []string, element string) []string {
	for i := 0; i < len(slice); {
		if slice[i] == element && i != len(slice)-1 {
			slice = append(slice[:i], slice[i+1:]...)
		} else if slice[i] == element && i == len(slice)-1 {
			slice = slice[:i]
		} else {
			i++
		}
	}
	return slice
}

func subProcess(args []string) *exec.Cmd {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[-] Error: %s\n", err)
	}
	return cmd
}

// BeDaemon convert the current process into a daemon process
func BeDaemon(arg string) {
	subProcess(stripSlice(os.Args, arg))
	fmt.Printf("[*] Daemon running in PID: %d PPID: %d\n", os.Getpid(), os.Getppid())
	os.Exit(0)
}
