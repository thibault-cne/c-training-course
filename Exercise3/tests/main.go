package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

var (
	executable = flag.String("executable", "../solution/build/phoneBook", "Provide the path to your executable. By default it will use the one from the solution.")
)

func main() {
	if _, err := exec.LookPath(*executable); err != nil {
		fmt.Printf("No executable found.\n")
		return
	}

	display(colorize("[STARTING TESTS]", Yellow), 0)

	runEndTests(endTests)

	display(colorize("[END TESTS]", Yellow), 0)
}

func runEndTests(tests []Test) {
	for _, t := range tests {
		display(colorize("[START "+strings.ToUpper(t.Name)+" TEST]", Blue), 1)

		for _, tSet := range t.TestSet {
			cmd, stdin, _ := startExecutable(*executable)

			for _, a := range tSet.StdinArgs {
				stdin.Write([]byte(a + "\n"))
			}

			if err := cmd.Wait(); err.Error() == tSet.Result {
				display(colorize("✓ Success : ", Green)+tSet.ItName, 2)
			} else {
				display(colorize("x Failed : ", Red)+tSet.ItName, 2)
				display(fmt.Sprintf("Expected %s but got %s instead.", tSet.Result, err.Error()), 3)
			}
		}

		display(colorize("[END "+strings.ToUpper(t.Name)+" TEST]", Blue), 1)
	}
}
