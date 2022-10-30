package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"exercise2.tests/colorize"
	"exercise2.tests/questions"
)

var (
	question = flag.Bool("questions", false, "Must be present to provide tests.")

	tests = []questions.Test{
		questions.T_questionOneOne,
		questions.T_questionTwoTwo,
		questions.T_questionTwoThree,
		questions.T_questionThreeTwo,
		questions.T_questionThreeThree,
	}
)

func main() {
	// We get the value of the question flag passed in the call of the executable
	flag.Parse()

	if !*question {
		fmt.Println(colorize.RedString("[ERROR]"), "Questions flag must be enabled to provide tests.")
		return
	}

	colorize.Yellow("[STARTING TESTS]\n")

	s := launchQuestions()

	colorize.Yellow("[END TESTS]\n")

	os.Exit(s)
}

func launchQuestions() int {
	var exitStatus int

	exitStatus = 0

	for _, t := range tests {

		if _, err := exec.LookPath(t.Executable); err == nil {
			res := questions.ExecuteTest(t)

			if res == 1 {
				exitStatus = 1
			}
		}
	}

	return exitStatus
}
