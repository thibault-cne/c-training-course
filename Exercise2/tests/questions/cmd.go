package questions

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func startExecutable(executable string, args []string) (*exec.Cmd, io.WriteCloser, *bufio.Scanner) {
	cmd := exec.Command("stdbuf", executableOptions(executable, args)...)

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

func executableOptions(executable string, args []string) []string {
	arr := []string{"-oL", executable}
	arr = append(arr, args...)

	return arr
}

func getLastLine(stdout *bufio.Scanner) string {
	var (
		previousLine string
		line         string
	)

	previousLine = stdout.Text()
	stdout.Scan()
	line = stdout.Text()

	for line != "" {
		previousLine = line
		stdout.Scan()
		line = stdout.Text()
	}

	return previousLine
}

func ExecuteTest(t Test) int {
	var res int

	res = 0

	display(colorize("[START "+strings.ToUpper(t.Name)+" TEST]", Blue), 1)

	for _, tSet := range t.TestSet {
		cmd, stdin, stdout := startExecutable(t.Executable, tSet.Args)

		for _, args := range tSet.StdinArgs {
			stdin.Write([]byte(args + "\n"))
		}

		answer := getLastLine(stdout)

		if answer != tSet.Result {
			res = 1
			display(colorize("x Failed : ", Red)+tSet.ItName, 2)
			display(fmt.Sprintf("Expected %s but got %s instead.", tSet.Result, answer), 3)
		} else {
			display(colorize("✓ Success : ", Green)+tSet.ItName, 2)
		}

		// We end the command
		cmd.Wait()

	}

	display(colorize("[END "+strings.ToUpper(t.Name)+" TEST]", Blue), 1)

	return res
}
