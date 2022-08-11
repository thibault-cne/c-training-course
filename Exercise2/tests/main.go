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
		fnc  func()
	}{
		{
			exec: "../build/questionOneOne.exe",
			fnc: func() {
				questions.QuestionOneOne()
			},
		},
		{
			exec: "../build/questionTwoTwo.exe",
			fnc: func() {
				questions.QuestionTwoTwo()
			},
		},
		{
			exec: "../build/questionTwoThree.exe",
			fnc: func() {
				questions.QuestionTwoThree()
			},
		},
		{
			exec: "../build/questionThreeTwo.exe",
			fnc: func() {
				questions.QuestionThreeTwo()
			},
		},
		{
			exec: "../build/questionThreeThree.exe",
			fnc: func() {
				questions.QuestionThreeThree()
			},
		},
	}

	for _, c := range cases {
		if _, err := exec.LookPath(c.exec); err == nil {
			c.fnc()
		}
	}
}
