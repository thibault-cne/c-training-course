package main

import (
	"bufio"
	"io"
	"os/exec"
)

func startExecutable(executable string) (*exec.Cmd, io.WriteCloser, *bufio.Scanner) {
	cmd := exec.Command("stdbuf", "-oL", executable)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	so, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	cmd.Start()

	stdout := bufio.NewScanner(so)

	return cmd, stdin, stdout
}
