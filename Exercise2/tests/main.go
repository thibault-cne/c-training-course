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
)

func main() {
	// We get the value of the question flag passed in the call of the executable
	// For exemple : tests -question=q1
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
	cases := []struct {
		exec string
		fnc  func(string)
	}{
		{
			exec: "../build/questionOneOne",
			fnc: func(exec string) {
				questions.QuestionOneOne(exec)
			},
		},
		{
			exec: "../build/questionTwoTwo",
			fnc: func(exec string) {
				questions.QuestionTwoTwo(exec)
			},
		},
		{
			exec: "../build/questionTwoThree",
			fnc: func(exec string) {
				questions.QuestionTwoThree(exec)
			},
		},
		{
			exec: "../build/questionThreeTwo",
			fnc: func(exec string) {
				questions.QuestionThreeTwo(exec)
			},
		},
		{
			exec: "../build/questionThreeThree",
			fnc: func(exec string) {
				questions.QuestionThreeThree(exec)
			},
		},
	}

	for _, c := range cases {
		if _, err := exec.LookPath(c.exec); err == nil {
			c.fnc(c.exec)
		}
	}
}
