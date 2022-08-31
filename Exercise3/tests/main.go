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

	Yellow("[START TESTS]\n")

	runEndTests(endTests)

	Yellow("[END TESTS]\n")
}

func runEndTests(tests []Test) {
	for _, t := range tests {
		Blue("	[START %s TEST]\n", strings.ToUpper(t.Name))

		for _, tSet := range t.TestSet {
			cmd, stdin, _ := startExecutable(*executable)

			for _, a := range tSet.StdinArgs {
				stdin.Write([]byte(a + "\n"))
			}

			if err := cmd.Wait(); err.Error() == tSet.Result {
				Green("		✓ Success : %s\n", tSet.ItName)
			} else {
				Red("		x Failed : %s\n", tSet.ItName)
				fmt.Printf("		Expected %s but got %s instead.\n", tSet.Result, err.Error())
			}
		}

		Blue("	[END %s TEST]\n", strings.ToUpper(t.Name))
	}
}
