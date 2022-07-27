package questions

import (
	"bufio"
	"io"
	"os/exec"
)

func startCommand(cmd *exec.Cmd) (io.WriteCloser, *bufio.Scanner) {
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

	return stdin, stdout
}
