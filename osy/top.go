package osy

import (
	"os"
	"os/exec"
)

func ExecTop() {
	command := exec.Command("top")
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}
