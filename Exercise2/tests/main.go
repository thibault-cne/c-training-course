package main

import (
	"flag"
	"fmt"
	"os/exec"

	"exercise2.tests/questions"
	"github.com/fatih/color"
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
		fmt.Println(color.RedString("[ERROR]"), "Questions flag must be enabled to provide tests.")
		return
	}

	fmt.Println(color.YellowString("[STARTING TESTS]"))

	launchQuestions()

	fmt.Println(color.YellowString("[END OF TESTS]"))
}

func launchQuestions() {
	for _, t := range tests {

		if _, err := exec.LookPath(t.Executable); err == nil {
			questions.ExecuteTest(t)
		}
	}
}
